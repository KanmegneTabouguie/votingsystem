// candidate/candidate_list.go

package candidate

// CandidateList represents a list of candidates
type CandidateList struct {
	Candidates []*Candidate
}

// NewCandidateList creates a new CandidateList
func NewCandidateList() *CandidateList {
	return &CandidateList{
		Candidates: []*Candidate{},
	}
}

// AddCandidate adds a new candidate to the list
func (cl *CandidateList) AddCandidate(candidate *Candidate) {
	cl.Candidates = append(cl.Candidates, candidate)
}

// GetCandidates returns the list of candidates
func (cl *CandidateList) GetCandidates() []*Candidate {
	return cl.Candidates
}

// GetCandidateByID returns a candidate by their ID
func (cl *CandidateList) GetCandidateByID(id int) *Candidate {
	for _, candidate := range cl.Candidates {
		if candidate.ID == id {
			return candidate
		}
	}
	return nil // Candidate with the given ID not found
}

// RemoveCandidate removes a candidate from the list by their ID
func (cl *CandidateList) RemoveCandidateByID(id int) {
	for i, candidate := range cl.Candidates {
		if candidate.ID == id {
			// Remove the candidate from the list
			cl.Candidates = append(cl.Candidates[:i], cl.Candidates[i+1:]...)
			return
		}
	}
}

// UpdateCandidate updates a candidate's information in the list
func (cl *CandidateList) UpdateCandidate(candidate *Candidate) {
	for i, existingCandidate := range cl.Candidates {
		if existingCandidate.ID == candidate.ID {
			// Update the candidate in the list
			cl.Candidates[i] = candidate
			return
		}
	}
	// Candidate not found, add them to the list
	cl.AddCandidate(candidate)
}
