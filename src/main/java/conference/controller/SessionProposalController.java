package conference.controller;

import conference.dto.SessionProposalRequest;
import conference.entity.SessionProposal;
import conference.entity.User;
import conference.service.SessionProposalService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;

@RestController
@RequestMapping("/api/session-proposals")
public class SessionProposalController {

    private final SessionProposalService sessionProposalService;

    @Autowired
    public SessionProposalController(SessionProposalService sessionProposalService) {
        this.sessionProposalService = sessionProposalService;
    }

    @PostMapping
    public ResponseEntity<SessionProposal> createProposal(@RequestBody SessionProposalRequest sessionProposalRequest) {
        try {
            SessionProposal proposal = sessionProposalService.createProposal(sessionProposalRequest.getTitle(), sessionProposalRequest.getDescription(), sessionProposalRequest.getStartTime(), sessionProposalRequest.getEndTime(), sessionProposalRequest.getUserId());
            return ResponseEntity.status(HttpStatus.CREATED).body(proposal);
        } catch (RuntimeException e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(null);
        }
    }

    @PutMapping("/{proposalId}")
    public ResponseEntity<String> editProposal(
            @PathVariable Long proposalId,
            @RequestParam String title,
            @RequestParam String description,
            @RequestParam LocalDateTime startTime,
            @RequestParam LocalDateTime endTime
    ) {
        boolean success = sessionProposalService.editProposal(title, description, startTime, endTime, proposalId);
        if (success) {
            return ResponseEntity.ok("Proposal updated successfully.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Proposal not found.");
        }
    }

    @DeleteMapping("/{proposalId}")
    public ResponseEntity<String> deleteProposal(@PathVariable Long proposalId) {
        boolean success = sessionProposalService.deleteProposal(proposalId);
        if (success) {
            return ResponseEntity.ok("Proposal deleted successfully.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Proposal not found.");
        }
    }

    @GetMapping
    public ResponseEntity<List<SessionProposal>> viewProposals() {
        List<SessionProposal> proposals = sessionProposalService.viewProposals();
        return ResponseEntity.ok(proposals);
    }

    @PutMapping("/accept/{proposalId}")
    public ResponseEntity<String> acceptProposal(@PathVariable("proposalId") Long proposalId) {
        boolean success = sessionProposalService.acceptProposals(proposalId);
        if (success) {
            return ResponseEntity.ok("Proposal accepted.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Proposal not found.");
        }
    }

    @PutMapping("/reject/{proposalId}")
    public ResponseEntity<String> rejectProposal(@PathVariable("proposalId") Long proposalId) {
        boolean success = sessionProposalService.rejectProposals(proposalId);
        if (success) {
            return ResponseEntity.ok("Proposal rejected.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Proposal not found.");
        }
    }
}
