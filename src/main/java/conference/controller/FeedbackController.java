package conference.controller;

import conference.dto.FeedbackRequest;
import conference.entity.Feedback;
import conference.service.FeedbackService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/feedback")
public class FeedbackController {

    private final FeedbackService feedbackService;

    @Autowired
    public FeedbackController(FeedbackService feedbackService) {
        this.feedbackService = feedbackService;
    }

    @PostMapping
    public ResponseEntity<Feedback> leaveFeedback(@RequestBody FeedbackRequest feedback) {
        try {
            Feedback savedFeedback = feedbackService.leaveFeedback(feedback.getSessionId(), feedback.getUserId(), feedback.getFeedback());
            return ResponseEntity.status(HttpStatus.CREATED).body(savedFeedback);
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).build();
        }
    }

    @DeleteMapping("/{feedbackId}")
    public ResponseEntity<String> removeFeedback(@PathVariable("feedbackId") Long feedbackId) {
        try {
            feedbackService.removeFeedback(feedbackId);
            return ResponseEntity.ok("Feedback removed successfully.");
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Feedback not found.");
        }
    }
}
