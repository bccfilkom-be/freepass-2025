package conference.entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@Table(name = "feedback")
public class Feedback {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "feedback_id")
    private Long id;

    @ManyToOne
    @JoinColumn(name = "session_id", nullable = false)
    @JsonIgnoreProperties({"hibernateLazyInitializer", "handler"})
    private Session session;

    @ManyToOne
    @JoinColumn(name = "user_id", nullable = false)
    @JsonIgnoreProperties({"hibernateLazyInitializer", "handler"})
    private User user;

    @Column(name = "feedback_text", nullable = false)
    private String feedbackText;

    @Column(name = "created_at", nullable = false)
    private LocalDateTime createdAt;

    public Feedback(Session session, User user, String feedbackText) {
        this.session = session;
        this.user = user;
        this.feedbackText = feedbackText;
        this.createdAt = LocalDateTime.now();
    }

    public Feedback() {
    }

    public Session getSession() {
        return session;
    }

    public void setSession(Session session) {
        this.session = session;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
