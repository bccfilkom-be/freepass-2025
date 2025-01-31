package conference.service;

import conference.entity.Feedback;
import conference.entity.Session;
import conference.repository.FeedbackRepository;
import conference.repository.SessionRepository;
import conference.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class FeedbackService {

    private final FeedbackRepository feedbackRepository;
    private final UserRepository userRepository;
    private final SessionRepository sesionRepository;

    @Autowired
    public FeedbackService(FeedbackRepository feedbackRepository, UserRepository userRepository, SessionRepository sesionRepository) {
        this.feedbackRepository = feedbackRepository;
        this.userRepository = userRepository;
        this.sesionRepository = sesionRepository;
    }

    public Feedback leaveFeedback(Long sessionId, Long userId, String feedback) {
        if (sesionRepository.existsById(sessionId) && sesionRepository.existsById(sessionId)) {
            return feedbackRepository.save(new Feedback(sesionRepository.getById(sessionId), userRepository.getById(userId), feedback));
        }
        return null;
    }

    public void removeFeedback(Long feedbackId) {
        feedbackRepository.deleteById(feedbackId);
    }
}
