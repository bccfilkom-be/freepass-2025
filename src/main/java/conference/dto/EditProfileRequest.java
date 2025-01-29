package conference.dto;

public class editProfileRequest {
    private String prevUsername;
    private String newUsername;
    private String newEmail;
    private String newFullName;

    public editProfileRequest(String prevUsername, String newUsername, String newEmail, String newFullName) {
        this.prevUsername = prevUsername;
        this.newUsername = newUsername;
        this.newEmail = newEmail;
        this.newFullName = newFullName;
    }
}
