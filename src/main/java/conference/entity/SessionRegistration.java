package conference.entity;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
@Table(name = "session_registrations")
public class SessionRegistration {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "registration_id")
    private Long id;

    @ManyToOne
    @JoinColumn(name = "user_id", nullable = false)
    private User user;

    @ManyToOne
    @JoinColumn(name = "session_id", nullable = false)
    private Session session;

    @Column(name = "registration_time", nullable = false)
    private LocalDateTime registrationTime;

    public Session getSession() {
        return session;
    }

    public SessionRegistration(User user, Session session) {
        this.user = user;
        this.session = session;
        this.registrationTime = LocalDateTime.now();
    }
}
