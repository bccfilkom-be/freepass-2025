package conference.repository;

import conference.entity.Session;
import jakarta.transaction.Transactional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.NativeQuery;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;

@Repository
public interface SessionRepository extends JpaRepository<Session, Long> {
    @Modifying
    @Transactional
    @NativeQuery("UPDATE sessions SET title = :new_title, description = :new_desc, start_time = :new_start, end_time = :new_end, max_seats = :new_max WHERE session_id = :id")
    int editSession(@Param("new_title") String title,
                         @Param("new_desc") String description,
                         @Param("new_start") LocalDateTime startTime,
                         @Param("new_end") LocalDateTime endTime,
                         @Param("new_max") Integer maxSeats,
                         @Param("id") Long sessionId);
}

