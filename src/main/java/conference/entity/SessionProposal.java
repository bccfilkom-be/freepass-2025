package conference.entity;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@Table(name = "session_proposals")
public class SessionProposal {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "proposal_id")
    private Long id;

    @Column(name = "title", nullable = false)
    private String title;

    @Column(name = "description")
    private String description;

    @Column(name = "start_time", nullable = false)
    private LocalDateTime startTime;

    @Column(name = "end_time", nullable = false)
    private LocalDateTime endTime;

    @ManyToOne
    @JoinColumn(name = "created_by", nullable = false)
    private User createdBy;

    @Column(name = "status", nullable = false)
    private String status;

    public SessionProposal(String title, String description, LocalDateTime startTime, LocalDateTime endTime, User createdBy, String status) {
        this.title = title;
        this.description = description;
        this.startTime = startTime;
        this.endTime = endTime;
        this.createdBy = createdBy;
        this.status = status;
    }

    public SessionProposal() {
    }

    public Long getId() {
        return id;
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

    public User getCreatedBy() {
        return createdBy;
    }

    public String getStatus() {
        return status;
    }
}
