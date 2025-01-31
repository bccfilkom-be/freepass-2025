package conference.dto;

import java.time.LocalDateTime;

public class SessionProposalRequest {
    private String title;
    private String description;
    private LocalDateTime startTime;
    private LocalDateTime endTime;
    private Long userId;

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

    public Long getUserId() {
        return userId;
    }
}
