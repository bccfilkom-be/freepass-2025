package conference.repository;

import conference.entity.SessionRegistration;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface SessionRegistrationRepository extends JpaRepository<SessionRegistration, Long> {
    boolean existsByUserIdAndSessionId(Long userId, Long sessionId);

    List<SessionRegistration> findByUserId(Long userId);
}
