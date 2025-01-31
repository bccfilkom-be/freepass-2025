package conference.controller;

import conference.dto.UserDto;
import conference.dto.EditProfileRequest;
import conference.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/users")
public class UserController {

    private final UserService userService;

    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping("/profile/{username}")
    public ResponseEntity<?> viewProfile(@PathVariable("username") String username) {
        UserDto userProfile = userService.viewProfile(username);
        if (userProfile != null) {
            return ResponseEntity.ok(userProfile);
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("User not found.");
        }
    }

    @PutMapping("/profile")
    public ResponseEntity<?> editProfile (@RequestBody EditProfileRequest editProfileRequest) {
        boolean updated = userService.editProfile(editProfileRequest.getPrevUsername(), editProfileRequest.getNewEmail(), editProfileRequest.getNewFullName(), editProfileRequest.getNewUsername());
        if (updated) {
            return ResponseEntity.ok("Profile updated successfully.");
        } else {
            return ResponseEntity.status(HttpStatus.CONFLICT).body("Email or username already exists or the user does not exist.");
        }
    }
}
