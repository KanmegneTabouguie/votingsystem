package main

import (
	"fmt"
	"voting-system/blockchain"
	"voting-system/candidate"
	"voting-system/voter"
)

func main() {
	// Initialize blockchain
	bc := blockchain.NewBlockchain()

	// Initialize candidate list
	candidateList := candidate.NewCandidateList()

	// Initialize voter list
	voterList := voter.NewVoterList()

	// Add sample candidates
	for i := 1; i <= 2; i++ {
		fmt.Printf("Enter the name of Candidate %d: ", i)
		var candidateName string
		fmt.Scan(&candidateName)
		candidateList.AddCandidate(&candidate.Candidate{ID: i, Name: candidateName, Description: ""})
	}

	// Display available candidates
	fmt.Println("\nAvailable Candidates:")
	for _, candidate := range candidateList.GetCandidates() {
		fmt.Printf("ID: %d, Name: %s\n", candidate.ID, candidate.Name)
	}

	// Add sample voters
	for i := 1; i <= 2; i++ {
		fmt.Printf("Enter the name of Voter %d: ", i)
		var voterName string
		fmt.Scan(&voterName)
		voterList.AddVoter(&voter.Voter{ID: i, Name: voterName, IsVoted: false})
	}

	// Simulate voting
	// Simulate voting
	for i := 1; i <= 2; i++ {
		var voterID, candidateID int

		fmt.Printf("\nVoter %d, enter your ID: ", i)
		fmt.Scan(&voterID)

		selectedVoter, err := voterList.GetVoterByID(voterID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if selectedVoter.IsVoted {
			fmt.Println("Error: Voter has already voted")
			return
		}

		fmt.Print("Enter the ID of the candidate you want to vote for: ")
		fmt.Scan(&candidateID)

		selectedCandidate := candidateList.GetCandidateByID(candidateID)
		if selectedCandidate == nil {
			fmt.Println("Candidate not found")
			return
		}

		// Perform voting
		transaction := &blockchain.Transaction{
			Sender:    fmt.Sprintf("Voter %d", selectedVoter.ID),
			Recipient: fmt.Sprintf("Candidate %d", selectedCandidate.ID),
			Amount:    1, // Voting amount or any value representing a vote
		}

		// Mine a new block with the voting transaction
		proofOfWork := blockchain.NewProofOfWork(2, "")
		proof, hash := proofOfWork.Run()

		// Print the hash to the console
		fmt.Println("Block Hash:", hash)
		// Add the block to the blockchain
		bc.AddBlock([]*blockchain.Transaction{transaction}, proof)

		// Update voter status to voted
		selectedVoter.IsVoted = true
		voterList.UpdateVoter(selectedVoter)
	}

	// Output the blockchain for verification
	fmt.Println("\nBlockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Hash: %s\n", block.CalculateHash())
	}

	// Display election results
	fmt.Println("\nElection Results:")
	for _, candidate := range candidateList.GetCandidates() {
		fmt.Printf("%s: %d votes\n", candidate.Name, bc.CountVotes(fmt.Sprintf("%d", candidate.ID)))
	}

	// Output the candidate list with campaign details
	fmt.Println("\nCandidate List:")
	for _, candidate := range candidateList.GetCandidates() {
		fmt.Printf("ID: %d, Name: %s, Description: %s\n", candidate.ID, candidate.Name, candidate.Description)
	}

	// Output the voter list
	fmt.Println("\nVoter List:")
	for _, voter := range voterList.GetVoters() {
		fmt.Printf("ID: %d, Name: %s, Voted: %t\n", voter.ID, voter.Name, voter.IsVoted)
	}
}
