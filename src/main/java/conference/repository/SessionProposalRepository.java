package conference.repository;

import conference.entity.SessionProposal;
import conference.entity.User;
import jakarta.transaction.Transactional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.NativeQuery;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;

@Repository
public interface SessionProposalRepository extends JpaRepository<SessionProposal, Long> {
    List<SessionProposal> findByCreatedBy(User createdBy);

    @Modifying
    @Transactional
    @NativeQuery("UPDATE session_proposals SET status = :new_status WHERE proposal_id = :id")
    int editStatus(@Param("new_status") String status, @Param("id") Long id);

    @Modifying
    @Transactional
    @NativeQuery("UPDATE session_proposals SET title = :new_title, description = :new_desc, start_time = :new_start, end_time = new_end WHERE proposal_id = :id")
    int editProposal(@Param("new_title") String title,
                        @Param("new_desc") String description,
                        @Param("new_start") LocalDateTime startTime,
                        @Param("new_end") LocalDateTime endTime,
                         @Param("id") Long proposalId);
}

