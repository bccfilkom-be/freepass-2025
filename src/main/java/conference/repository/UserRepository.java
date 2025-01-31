package conference.repository;

import conference.entity.User;
import jakarta.transaction.Transactional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.NativeQuery;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByUsername(String username);
    boolean existsByUsername(String username);
    boolean existsByEmail(String email);

    @Modifying
    @Transactional
    @NativeQuery("UPDATE users SET username = :new_username, email = :new_email, full_name = :new_fullname WHERE username = :prev_username")
    int editProfile(@Param("new_username") String username,
                        @Param("new_email") String email,
                        @Param("new_fullname") String fullName,
                        @Param("prev_username") String userID);

    @Modifying
    @Transactional
    @NativeQuery("UPDATE users SET role_id = 2 WHERE user_id = :id")
    int editRole(@Param("id") Long id);
}
