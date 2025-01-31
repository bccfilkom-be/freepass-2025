package conference.controller;

import conference.dto.RegisterForSessionRequest;
import conference.entity.SessionRegistration;
import conference.service.SessionRegistrationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/session-registrations")
public class SessionRegistrationController {

    private final SessionRegistrationService sessionRegistrationService;

    @Autowired
    public SessionRegistrationController(SessionRegistrationService sessionRegistrationService) {
        this.sessionRegistrationService = sessionRegistrationService;
    }

    @PostMapping
    public ResponseEntity<?> registerForSession(@RequestBody RegisterForSessionRequest registerForSessionRequest) {
        try {
            SessionRegistration registration = sessionRegistrationService.registerForSession(registerForSessionRequest.getUserId(), registerForSessionRequest.getSessionId());
            return ResponseEntity.status(HttpStatus.CREATED).body(registration);
        } catch (RuntimeException e) {
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(e.getMessage());
        }
    }
}
