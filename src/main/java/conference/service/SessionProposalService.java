package conference.service;

import conference.entity.SessionProposal;
import conference.entity.User;
import conference.repository.SessionProposalRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;

@Service
public class SessionProposalService {
    private final SessionProposalRepository sessionProposalRepository;

    @Autowired
    public SessionProposalService(SessionProposalRepository sessionProposalRepository) {
        this.sessionProposalRepository = sessionProposalRepository;
    }

    public SessionProposal createProposal(String title, String description, LocalDateTime startTime, LocalDateTime endTime, User createdBy){
        List<SessionProposal> list = sessionProposalRepository.findByCreatedBy(createdBy);
        if (!list.isEmpty()){
            for (SessionProposal s : list){
                if (startTime.isAfter(s.getEndTime()) || endTime.isBefore(s.getStartTime())) {
                    continue;
                }
                throw new RuntimeException("The session time period is already proposed!");
            }
        }
        return sessionProposalRepository.save(new SessionProposal(title, description, startTime, endTime, createdBy, "Pending"));
    }

    public boolean editProposal(String title, String description, LocalDateTime startTime, LocalDateTime endTime, Long proposalId){
        return sessionProposalRepository.editProposal(title, description, startTime, endTime, proposalId);
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
            sessionProposalRepository.editStatus("Accepted");
            return true;
        } else {
            return false;
        }
    }

    public boolean rejectProposals(Long proposalId){
        if (sessionProposalRepository.existsById(proposalId)){
            sessionProposalRepository.editStatus("Rejected");
            return true;
        } else {
            return false;
        }
    }
}
