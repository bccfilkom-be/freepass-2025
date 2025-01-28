package conference.service;

import conference.entity.Feedback;
import conference.repository.FeedbackRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class FeedbackService {

    private final FeedbackRepository feedbackRepository;

    @Autowired
    public FeedbackService(FeedbackRepository feedbackRepository) {
        this.feedbackRepository = feedbackRepository;
    }

    public Feedback leaveFeedback(Feedback feedback) {
        return feedbackRepository.save(feedback);
    }

    public void removeFeedback(Long feedbackId) {
        feedbackRepository.deleteById(feedbackId);
    }
}
