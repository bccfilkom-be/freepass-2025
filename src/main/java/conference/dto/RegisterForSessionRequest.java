package conference.dto;

public class RegisterForSessionRequest {
    private Long userId;
    private Long sessionId;

    public Long getUserId() {
        return userId;
    }

    public Long getSessionId() {
        return sessionId;
    }
}
