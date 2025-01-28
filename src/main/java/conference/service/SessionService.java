package conference.service;

import conference.entity.Session;
import conference.repository.SessionRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
public class SessionService {

    private final SessionRepository sessionRepository;

    @Autowired
    public SessionService(SessionRepository sessionRepository) {
        this.sessionRepository = sessionRepository;
    }

    // View all sessions
    public List<Session> viewAllSessions() {
        return sessionRepository.findAll();
    }

    public boolean editSession(Long sessionId, String title, String description, LocalDateTime startTime, LocalDateTime endTime, Integer maxSeats){
        if (sessionRepository.existsById(sessionId)){
            return false;
        } else {
            sessionRepository.editSession(title, description, startTime, endTime, maxSeats, sessionId);
            return true;
        }
    }

    public boolean deleteSession(Long sessionId){
        if (sessionRepository.existsById(sessionId)){
            sessionRepository.deleteById(sessionId);
            return true;
        } else {
            return false;
        }
    }
}
