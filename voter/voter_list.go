// voter/voter_list.go

package voter

import (
	"errors"
)

// VoterList represents a list of voters
type VoterList struct {
	Voters []*Voter
}

// NewVoterList creates a new VoterList
func NewVoterList() *VoterList {
	return &VoterList{
		Voters: []*Voter{},
	}
}

// AddVoter adds a new voter to the list
func (vl *VoterList) AddVoter(voter *Voter) {
	vl.Voters = append(vl.Voters, voter)
}

// GetVoters returns the list of voters
func (vl *VoterList) GetVoters() []*Voter {
	return vl.Voters
}

// GetVoterByID returns a voter by their ID
func (vl *VoterList) GetVoterByID(id int) (*Voter, error) {
	for _, voter := range vl.Voters {
		if voter.ID == id {
			return voter, nil
		}
	}
	return nil, errors.New("Voter not found")
}

// RemoveVoterByID removes a voter from the list by their ID
func (vl *VoterList) RemoveVoterByID(id int) error {
	for i, voter := range vl.Voters {
		if voter.ID == id {
			// Remove the voter from the list
			vl.Voters = append(vl.Voters[:i], vl.Voters[i+1:]...)
			return nil
		}
	}
	return errors.New("Voter not found")
}

// UpdateVoter updates a voter's information in the list
func (vl *VoterList) UpdateVoter(voter *Voter) error {
	for i, existingVoter := range vl.Voters {
		if existingVoter.ID == voter.ID {
			// Update the voter in the list
			vl.Voters[i] = voter
			return nil
		}
	}
	return errors.New("Voter not found, cannot update")
}
