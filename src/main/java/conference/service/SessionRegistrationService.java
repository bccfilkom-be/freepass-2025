package conference.service;

import conference.entity.Session;
import conference.entity.SessionRegistration;
import conference.entity.User;
import conference.repository.SessionRegistrationRepository;
import conference.repository.SessionRepository;
import conference.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Service
public class SessionRegistrationService {

    private final SessionRegistrationRepository registrationRepository;
    private final SessionRepository sessionRepository;
    private final UserRepository userRepository;

    @Autowired
    public SessionRegistrationService(SessionRegistrationRepository registrationRepository, SessionRepository sessionRepository, UserRepository userRepository) {
        this.registrationRepository = registrationRepository;
        this.sessionRepository = sessionRepository;
        this.userRepository = userRepository;
    }

    public SessionRegistration registerForSession(Long userId, Long sessionId) {
        Optional<Session> tempSess = sessionRepository.findById(sessionId);
        Optional<User> tempUser = userRepository.findById(userId);
        Session sess;
        User user;
        if (tempSess.isEmpty()){
            throw new RuntimeException("There is no session!");
        } else {
            sess = tempSess.get();
            if (LocalDateTime.now().isAfter(sess.getStartTime())){
                throw new RuntimeException("Session already started!");
            }
        }
        if (tempUser.isEmpty()){
            throw new RuntimeException("There is no user!");
        } else {
            user = tempUser.get();
        }
        if (registrationRepository.existsByUserIdAndSessionId(userId, sessionId)){
            throw new RuntimeException("User is already registered for this session");
        } else {
            List<SessionRegistration> sr = registrationRepository.findByUserId(userId);
            for (SessionRegistration cur : sr){
                if (sess.getStartTime().isAfter(cur.getSession().getEndTime()) || sess.getEndTime().isBefore(cur.getSession().getStartTime())) {
                    continue;
                }
                throw new RuntimeException("The session time period is already registered!");
            }
        }

        return registrationRepository.save(new SessionRegistration(user, sess));
    }

}
