package conference.controller;

import conference.dto.EditSessionRequest;
import conference.entity.Session;
import conference.service.SessionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.util.List;

@RestController
@RequestMapping("/api/sessions")
public class SessionController {

    private final SessionService sessionService;

    @Autowired
    public SessionController(SessionService sessionService) {
        this.sessionService = sessionService;
    }

    @GetMapping
    public ResponseEntity<List<Session>> getAllSessions() {
        List<Session> sessions = sessionService.viewAllSessions();
        return ResponseEntity.ok(sessions);
    }

    @PutMapping
    public ResponseEntity<?> editSession(@RequestBody EditSessionRequest editSessionRequest) {
        boolean updated = sessionService.editSession(editSessionRequest.getSessionId(), editSessionRequest.getTitle(), editSessionRequest.getDescription(), editSessionRequest.getStartTime(), editSessionRequest.getEndTime(), editSessionRequest.getMaxSeats());
        if (updated) {
            return ResponseEntity.ok("Session updated successfully.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Session not found.");
        }
    }

    @DeleteMapping("/{sessionId}")
    public ResponseEntity<?> deleteSession(@PathVariable("sessionId") Long sessionId) {
        boolean deleted = sessionService.deleteSession(sessionId);
        if (deleted) {
            return ResponseEntity.ok("Session deleted successfully.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Session not found.");
        }
    }
}
