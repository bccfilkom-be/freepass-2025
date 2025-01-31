package conference.dto;

import java.time.LocalDateTime;

public class EditSessionRequest {
    private Long sessionId;
    private String title;
    private String description;
    private LocalDateTime startTime;
    private LocalDateTime endTime;
    private Integer maxSeats;

    public Long getSessionId() {
        return sessionId;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }

    public LocalDateTime getStartTime() {
        return startTime;
    }

    public LocalDateTime getEndTime() {
        return endTime;
    }

    public Integer getMaxSeats() {
        return maxSeats;
    }
}
