package conference.service;

import conference.entity.SessionProposal;
import conference.entity.User;
import conference.repository.SessionProposalRepository;
import conference.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

@Service
public class SessionProposalService {
    private final SessionProposalRepository sessionProposalRepository;
    private final UserRepository userRepository;

    @Autowired
    public SessionProposalService(SessionProposalRepository sessionProposalRepository, UserRepository userRepository) {
        this.sessionProposalRepository = sessionProposalRepository;
        this.userRepository = userRepository;
    }

    public SessionProposal createProposal(String title, String description, LocalDateTime startTime, LocalDateTime endTime, Long id){
        Optional<User> temp = userRepository.findById(id);
        User createdBy;
        if (temp.isPresent()){
            createdBy = temp.get();
        } else {
            return null;
        }
        if (startTime.isAfter(endTime)){
            return null;
        }
        List<SessionProposal> list = sessionProposalRepository.findByCreatedBy(createdBy);
        if (!list.isEmpty()){
            for (SessionProposal s : list){
                if (startTime.isAfter(s.getEndTime()) || endTime.isBefore(s.getStartTime())) {
                    continue;
                }
                return null;
            }
        }
        return sessionProposalRepository.save(new SessionProposal(title, description, startTime, endTime, createdBy, "Pending"));
    }

    public boolean editProposal(String title, String description, LocalDateTime startTime, LocalDateTime endTime, Long proposalId){
        if (sessionProposalRepository.existsById(proposalId) || endTime.isBefore(startTime)){
            int val = sessionProposalRepository.editProposal(title, description, startTime, endTime, proposalId);
            return val == 1;
        } else {
            return false;
        }
    }

    public boolean deleteProposal(Long proposalId){
        if (sessionProposalRepository.existsById(proposalId)){
            sessionProposalRepository.deleteById(proposalId);
            return true;
        }
        return false;
    }

    public List<SessionProposal> viewProposals(){
        return sessionProposalRepository.findAll();
    }

    public boolean acceptProposals(Long proposalId){
        if (sessionProposalRepository.existsById(proposalId)){
            sessionProposalRepository.editStatus("Accepted", proposalId);
            return true;
        } else {
            return false;
        }
    }

    public boolean rejectProposals(Long proposalId){
        if (sessionProposalRepository.existsById(proposalId)){
            sessionProposalRepository.editStatus("Rejected", proposalId);
            return true;
        } else {
            return false;
        }
    }
}
