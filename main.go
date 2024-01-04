package main

import (
	"fmt"
)

// FilmDistributionSystem represents the film distribution system
type FilmDistributionSystem struct {
	Permissions map[string]DistributorPermissions
}

// DistributorPermissions represents the permissions for a distributor
type DistributorPermissions struct {
	Include map[string]struct{}
	Exclude map[string]struct{}
}

// NewFilmDistributionSystem creates a new FilmDistributionSystem
func NewFilmDistributionSystem() *FilmDistributionSystem {
	return &FilmDistributionSystem{
		Permissions: make(map[string]DistributorPermissions),
	}
}

// AddPermissions adds permissions for a distributor
func (fds *FilmDistributionSystem) AddPermissions(distributor string, include, exclude string) {
	if _, exists := fds.Permissions[distributor]; !exists {
		fds.Permissions[distributor] = DistributorPermissions{
			Include: make(map[string]struct{}),
			Exclude: make(map[string]struct{}),
		}
	}

	if include != "" {
		fds.Permissions[distributor].Include[include] = struct{}{}
	}
	if exclude != "" {
		fds.Permissions[distributor].Exclude[exclude] = struct{}{}
	}
}

// CheckPermission checks if a distributor has permission for a given region
func (fds *FilmDistributionSystem) CheckPermission(distributor, region string) bool {
	permissions, exists := fds.Permissions[distributor]
	if !exists {
		return false
	}

	_, included := permissions.Include[region]
	_, excluded := permissions.Exclude[region]

	return included && !excluded
}

func main() {
	// Example usage:

	filmSystem := NewFilmDistributionSystem()

	// Define permissions for DISTRIBUTOR1
	filmSystem.AddPermissions("DISTRIBUTOR1", "INDIA", "")
	filmSystem.AddPermissions("DISTRIBUTOR1", "UNITEDSTATES", "")
	filmSystem.AddPermissions("DISTRIBUTOR1", "", "KARNATAKA-INDIA")
	filmSystem.AddPermissions("DISTRIBUTOR1", "", "CHENNAI-TAMILNADU-INDIA")

	// Define permissions for DISTRIBUTOR2, a subset of DISTRIBUTOR1
	filmSystem.AddPermissions("DISTRIBUTOR2", "INDIA", "")
	filmSystem.AddPermissions("DISTRIBUTOR2", "", "TAMILNADU-INDIA")

	// Define permissions for DISTRIBUTOR3, a subset of DISTRIBUTOR2
	filmSystem.AddPermissions("DISTRIBUTOR3", "HUBLI-KARNATAKA-INDIA", "")

	// Check permissions for specific regions
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR1", "CHICAGO-ILLINOIS-UNITEDSTATES")) // Should print true
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR1", "CHENNAI-TAMILNADU-INDIA"))       // Should print false
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR1", "BANGALORE-KARNATAKA-INDIA"))     // Should print false

	// Check permissions for DISTRIBUTOR2
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR2", "CHICAGO-ILLINOIS-UNITEDSTATES"))   // Should print true
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR2", "TIRUCHIRAPPALLI-TAMILNADU-INDIA")) // Should print false

	// Check permissions for DISTRIBUTOR3
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR3", "HUBLI-KARNATAKA-INDIA"))     // Should print true
	fmt.Println(filmSystem.CheckPermission("DISTRIBUTOR3", "BANGALORE-KARNATAKA-INDIA")) // Should print false
}
