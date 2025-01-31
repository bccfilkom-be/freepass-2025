package conference.dto;

public class FeedbackRequest {
    private Long sessionId;
    private Long userId;
    private String feedback;

    public Long getSessionId() {
        return sessionId;
    }

    public Long getUserId() {
        return userId;
    }

    public String getFeedback() {
        return feedback;
    }
}
