package conference.dto;

public class EditProfileRequest {
    private String prevUsername;
    private String newUsername;
    private String newEmail;
    private String newFullName;

    public String getPrevUsername() {
        return prevUsername;
    }

    public void setPrevUsername(String prevUsername) {
        this.prevUsername = prevUsername;
    }

    public String getNewUsername() {
        return newUsername;
    }

    public void setNewUsername(String newUsername) {
        this.newUsername = newUsername;
    }

    public String getNewEmail() {
        return newEmail;
    }

    public void setNewEmail(String newEmail) {
        this.newEmail = newEmail;
    }

    public String getNewFullName() {
        return newFullName;
    }

    public void setNewFullName(String newFullName) {
        this.newFullName = newFullName;
    }
}
