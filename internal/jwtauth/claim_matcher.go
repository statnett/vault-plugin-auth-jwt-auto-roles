package jwtauth

import "maps"

// node enables checking if claims satisfies bound claims in an indexed way. It
// is done by only traversing children through key-value pairs existing in the
// checked claims. E.g. only traverse to children["project_path"]["bar/baz"] if
// claims contains "project_path": "bar/baz". The claim satisfies the role names
// of all nodes' roleNames reached in traversion.
type node struct {
	// roleNames is a list of role names with all their bound claim's key-value
	// pairs checked between the root and this node. E.g. "foo" in root's
	// children["project_path"]["bar/baz"].children["ref"]["main"].roleNames
	// only if role "foo" has exactly two bound claims, "project_path" and
	// "ref", and the bound claim values contain at least "bar/baz" and "main",
	// respectively.
	roleNames []string
	// children maps bound claim key => claim value => child node. The same role
	// name can be present in multiple children's roleNames, but only due to the
	// role having multiple bound values for a claim key.
	children map[string]map[any]*node
}

// findBoundRoleNames traverses the tree to find the names of all roles a set of
// claims satisfies. It is done by collecting the role names of all nodes that
// can be reached by indexing children with the claim's value for each of the
// children map's (bound claim) keys. Note: only if claims contains the key.
func (n *node) findBoundRoleNames(claims map[string]any) []string {
	roleNames := n.roleNames
	for claimKey, children := range n.children {
		if claimValue, ok := claims[claimKey]; ok {
			if child := children[claimValue]; child != nil {
				roleNames = append(roleNames, child.findBoundRoleNames(claims)...)
			}
		}
	}
	return roleNames
}

type boundClaims map[string][]string

type roles map[string]boundClaims

// createMatchTree builds a tree able to match claims, based on a list of roles.
// It does so by recursively creating subtrees with subsets of the roles.
func (r roles) createMatchTree() *node {
	var roleNames []string
	// Any role with no bound claims left is satisfied at this node in the tree.
	// Extract the role name and remove the role.
	for roleName, boundClaims := range r {
		if len(boundClaims) == 0 {
			delete(r, roleName)
			roleNames = append(roleNames, roleName)
		}
	}

	return &node{
		children:  r.children(),
		roleNames: roleNames,
	}
}

// children iteratively creates new subtrees with all roles whose bound claims
// contains a selected claim key, although with that key removed from their
// bound claims.
//
// Each iteration adds a claim key to the children map, mapping bound values to
// other nodes. It acts as a verification check of that bound claim when
// checking claims in findBoundRoleNames. Not all claim keys are present in
// children, as to not create unnessesary permutations of the same set of bound
// claims. E.g. only one of
// children["project_path"]["bar/baz"].children["ref"]["main"] and
// children["ref"]["main"].children["project_path"]["bar/baz"] will exist.
func (r roles) children() map[string]map[any]*node {
	var children map[string]map[any]*node
	for len(r) > 0 {
		claimKey := r.anyClaimKey()

		groups := make(map[any]roles)
		for roleName, boundClaims := range r {
			if values, ok := boundClaims[claimKey]; ok {
				delete(r, roleName)
				delete(boundClaims, claimKey)
				for _, value := range values {
					if groups[value] == nil {
						groups[value] = make(roles)
					}
					groups[value][roleName] = maps.Clone(boundClaims)
				}
			}
		}

		if children == nil {
			children = make(map[string]map[any]*node)
		}
		children[claimKey] = make(map[any]*node)
		for value, roles := range groups {
			children[claimKey][value] = roles.createMatchTree()
		}
	}
	return children
}

func (r roles) anyClaimKey() string {
	for _, bc := range r {
		for k := range bc {
			return k
		}
	}
	panic("no roles have bound claims")
}
