package conference.service;

import conference.dto.UserDto;
import conference.entity.User;
import conference.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Objects;
import java.util.Optional;

@Service
public class UserService {

    private final UserRepository userRepository;

    @Autowired
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public boolean editProfile(String prevUsername,  String new_email, String new_fullname , String new_username){
        Optional<User> temp = userRepository.findByUsername(prevUsername);
        if (temp.isPresent()){
            User cur = temp.get();
            if ((userRepository.existsByEmail(new_email) && !Objects.equals(cur.getEmail(), new_email)) || (userRepository.existsByUsername(new_username)  && !Objects.equals(cur.getUsername(), new_email))){
                return false;
            } else {
                int val = userRepository.editProfile(new_username, new_email, new_fullname, prevUsername);
                return val == 1;
            }
        }
        return false;
    }

    public UserDto viewProfile(String username){
        Optional<User> temp = userRepository.findByUsername(username);
        if (temp.isPresent()){
            User cur = temp.get();
            return new UserDto(cur.getUsername(), cur.getEmail(), cur.getFullName());
        } else {
            return null;
        }
    }
}
