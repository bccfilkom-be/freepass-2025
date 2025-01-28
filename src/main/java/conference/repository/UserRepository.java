package conference.repository;

import conference.entity.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.NativeQuery;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import java.util.Optional;

@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByUsername(String username);
    boolean existsByUsername(String username);
    boolean existsByEmail(String email);

    @NativeQuery("UPDATE users SET username = :new_username, email = :new_email, full_name = :new_fullname WHERE user_id = :id")
    int editProfile(@Param("new_username") String username,
                        @Param("new_email") String email,
                        @Param("new_fullname") String fullName,
                        @Param("id") Long userID);

}
