package conference.controller;

import conference.service.AdminService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/admin")
public class AdminController {

    private final AdminService adminService;

    @Autowired
    public AdminController(AdminService adminService) {
        this.adminService = adminService;
    }

    @PostMapping("/coordinator/{id}")
    public ResponseEntity<String> addEventCoordinator(@PathVariable("id") Long id){
        boolean added = adminService.addEventCoordinator(id);
        if (added){
            return ResponseEntity.ok("Event Coordinator added.");
        } else {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("User not found.");
        }
    }

    @DeleteMapping("/remove/{id}")
    public ResponseEntity<String> removeFeedback(@PathVariable("id") Long id) {
        try {
            adminService.removeUser(id);
            return ResponseEntity.ok("User removed successfully.");
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("User not found.");
        }
    }
}
