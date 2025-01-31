package conference.service;

import conference.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AdminService {

    private final UserRepository userRepository;

    @Autowired
    public AdminService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public boolean addEventCoordinator(Long userId){
        if (userRepository.existsById(userId)){
            int val = userRepository.editRole(userId);
            return val == 1;
        } else {
            return false;
        }
    }

    public boolean removeUser(Long userId){
        if (userRepository.existsById(userId)){
            userRepository.deleteById(userId);
            return true;
        } else {
            return false;
        }
    }
}
