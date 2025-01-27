CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE OR REPLACE FUNCTION generate_ulid_at_time(target_time TIMESTAMP WITH TIME ZONE)
    RETURNS UUID AS
$$
DECLARE
    timestamp_ms BIGINT;
    rand_bytes   BYTEA;
    result       UUID;
BEGIN
    timestamp_ms := FLOOR(EXTRACT(EPOCH FROM target_time) * 1000);
    rand_bytes := gen_random_bytes(10);
    result := CONCAT_WS('-',
                        LPAD(TO_HEX(timestamp_ms >> 16), 8, '0'),
                        LPAD(TO_HEX((timestamp_ms & x'FFFF'::int)), 4, '0'),
                        LPAD(TO_HEX(x'4000'::int | (get_byte(rand_bytes, 0) & x'0FFF'::int)), 4, '0'),
                        LPAD(TO_HEX(x'8000'::int | (get_byte(rand_bytes, 1) & x'3FFF'::int)), 4, '0'),
                        CONCAT(
                                LPAD(TO_HEX(get_byte(rand_bytes, 2)), 2, '0'),
                                LPAD(TO_HEX(get_byte(rand_bytes, 3)), 2, '0'),
                                LPAD(TO_HEX(get_byte(rand_bytes, 4)), 2, '0'),
                                LPAD(TO_HEX(get_byte(rand_bytes, 5)), 2, '0'),
                                LPAD(TO_HEX(get_byte(rand_bytes, 6)), 2, '0'),
                                LPAD(TO_HEX(get_byte(rand_bytes, 7)), 2, '0')
                        )
              )::UUID;
    RETURN result;
END;
$$ LANGUAGE plpgsql;

DO
$$
    DECLARE
        -- User IDs declarations
        user1_id  UUID;
        user2_id  UUID;
        user3_id  UUID;
        user4_id  UUID;
        ec1_id    UUID;
        ec2_id    UUID;
        admin1_id UUID;
    BEGIN
        -- Initialize UUIDs
        user1_id := generate_ulid_at_time(NOW() - INTERVAL '30 days');
        user2_id := generate_ulid_at_time(NOW() - INTERVAL '30 days' + INTERVAL '1 hour');
        user3_id := generate_ulid_at_time(NOW() - INTERVAL '30 days' + INTERVAL '2 hours');
        user4_id := generate_ulid_at_time(NOW() - INTERVAL '30 days' + INTERVAL '3 hours');
        ec1_id := generate_ulid_at_time(NOW() - INTERVAL '30 days' + INTERVAL '4 hours');
        ec2_id := generate_ulid_at_time(NOW() - INTERVAL '30 days' + INTERVAL '5 hours');
        admin1_id := generate_ulid_at_time(NOW() - INTERVAL '30 days' + INTERVAL '6 hours');

        -- Users seeder
        INSERT INTO users (id, name, email, password_hash, role, bio, created_at)
        VALUES
            -- Regular users
            (user1_id, 'User 1', 'user1@seeder.nathakusuma.com',
             '$2a$12$Bvp.gDd3bNy.66etGgJsFe.KSJ6KmsxM/NWKA4BgUjI3WwHuhKHRS', 'user',
             'Software Engineer with 5 years experience', NOW() - INTERVAL '30 days'),
            (user2_id, 'User 2', 'user2@seeder.nathakusuma.com',
             '$2a$12$ygPrGPgl.icv/EDn3GxXlOD46osqFLAi4AE.PnVhByQkGkxo1pCQm', 'user', 'Product Manager at Tech Corp',
             NOW() - INTERVAL '30 days' + INTERVAL '1 hour'),
            (user3_id, 'User 3', 'user3@seeder.nathakusuma.com',
             '$2a$12$TpyW400ggb7OYJp4q3Zrou7XNN0G4NgobzClJYmLoXjgoMWsoKxa.', 'user', 'Full Stack Developer',
             NOW() - INTERVAL '30 days' + INTERVAL '2 hours'),
            (user4_id, 'User 4', 'user4@seeder.nathakusuma.com',
             '$2a$12$CPr7.LBs/.d.i8/f0HYjYOQkH9h9gnYIKxVM3WbCD40lUFnys5z9S', 'user', 'Back End Developer',
             NOW() - INTERVAL '30 days' + INTERVAL '3 hours'),
            -- Event coordinator
            (ec1_id, 'Event Coordinator 1', 'ec1@seeder.nathakusuma.com',
             '$2a$12$u6RpbKs906IT4Yt39px0medvda/vklB5kCfp0iPCnm1q9bek0Pl0u', 'event_coordinator',
             'Professional Event Coordinator', NOW() - INTERVAL '30 days' + INTERVAL '4 hours'),
            (ec2_id, 'Event Coordinator 2', 'ec2@seeder.nathakusuma.com',
             '$2a$12$9dDdCZFJs.hAR1TApoyYxOSPfYCqv.w.B2nuEfVFBtiItUitlv6du', 'event_coordinator',
             'Professional Event Coordinator', NOW() - INTERVAL '30 days' + INTERVAL '5 hours'),
            -- Admin
            (admin1_id, 'Admin User 1', 'admin1@seeder.nathakusuma.com',
             '$2a$12$TTXysJS9gxY/o.ODh5QTpupxGtzuSLQWN8jAkSKChq86/uOhvw1UO', 'admin', 'System Administrator',
             NOW() - INTERVAL '30 days' + INTERVAL '6 hours');

        -- Conferences seeder
        INSERT INTO conferences (id, title, description, speaker_name, speaker_title, target_audience, prerequisites,
                                 seats, starts_at, ends_at, host_id, status, created_at)
        VALUES
            -- Past conferences (approved)
            (generate_ulid_at_time(NOW() - INTERVAL '12 days'), 'Past Conference 1',
             'Description for past conference 1', 'Dr. Smith', 'Professor', 'Developers', 'Basic programming', 100,
             NOW() - INTERVAL '7 days', NOW() - INTERVAL '7 days' + INTERVAL '2 hours', user1_id, 'approved',
             NOW() - INTERVAL '12 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '12 days'), 'Past Conference 2',
             'Description for past conference 2', 'Jane Doe', 'Tech Lead', 'Architects', 'System design experience', 50,
             NOW() - INTERVAL '6 days', NOW() - INTERVAL '6 days' + INTERVAL '2 hours', user2_id, 'approved',
             NOW() - INTERVAL '12 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '17 days'), 'Past Conference 3', 'Advanced JavaScript Patterns',
             'Lisa Johnson', 'Senior JS Developer', 'Advanced developers', 'JavaScript experience', 120,
             NOW() - INTERVAL '9 days', NOW() - INTERVAL '9 days' + INTERVAL '2 hours', user1_id, 'approved',
             NOW() - INTERVAL '17 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '17 days'), 'Past Conference 4', 'Microservices Architecture',
             'Mike Chen', 'Solutions Architect', 'System architects', 'Distributed systems knowledge', 80,
             NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days' + INTERVAL '2 hours', user2_id, 'approved',
             NOW() - INTERVAL '17 days'),

            -- Current approved conference
            (generate_ulid_at_time(NOW() - INTERVAL '7 days'), 'Current Active Conference',
             'Currently running conference', 'Dr. Johnson', 'CTO', 'All developers', NULL, 200, NOW(),
             NOW() + INTERVAL '4 hours', user3_id, 'approved', NOW() - INTERVAL '7 days'),

            -- Future conferences (approved)
            (generate_ulid_at_time(NOW() - INTERVAL '7 days'), 'Future Conference 1',
             'Description for future conference 1', 'Alice Brown', 'Senior Developer', 'Junior developers', 'None', 150,
             NOW() + INTERVAL '5 days', NOW() + INTERVAL '5 days' + INTERVAL '2 hours', user1_id, 'approved',
             NOW() - INTERVAL '7 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '7 days'), 'Future Conference 2',
             'Description for future conference 2', 'Bob Williams', 'Architect', 'Senior developers',
             'Advanced programming', 75, NOW() + INTERVAL '9 days', NOW() + INTERVAL '9 days' + INTERVAL '2 hours',
             user2_id, 'approved', NOW() - INTERVAL '7 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Future Conference 3',
             'Description for future conference 3', 'Henry Ford', 'Tech Lead', 'All levels', NULL, 200,
             NOW() + INTERVAL '37 days', NOW() + INTERVAL '37 days' + INTERVAL '2 hours', user3_id, 'approved',
             NOW() - INTERVAL '3 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Future Conference 4',
             'Description for future conference 4', 'Ivy Chen', 'Senior Architect', 'Senior developers',
             'Architecture experience', 100, NOW() + INTERVAL '42 days',
             NOW() + INTERVAL '42 days' + INTERVAL '2 hours', user1_id, 'approved', NOW() - INTERVAL '3 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Future Conference 5', 'Cloud Native Applications',
             'Nathan Black', 'Cloud Architect', 'DevOps engineers', 'Kubernetes basics', 150,
             NOW() + INTERVAL '57 days', NOW() + INTERVAL '57 days' + INTERVAL '2 hours', user2_id, 'approved',
             NOW() - INTERVAL '3 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Future Conference 6', 'AI in Production',
             'Olivia Green', 'ML Engineer', 'Data scientists', 'Python, ML basics', 100, NOW() + INTERVAL '64 days',
             NOW() + INTERVAL '64 days' + INTERVAL '2 hours', user3_id, 'approved', NOW() - INTERVAL '3 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Future Conference 7', 'Blockchain Development',
             'Peter White', 'Blockchain Developer', 'Developers', 'Cryptography basics', 120,
             NOW() + INTERVAL '68 days', NOW() + INTERVAL '68 days' + INTERVAL '2 hours', user1_id, 'approved',
             NOW() - INTERVAL '3 days'),

            -- Pending conferences (one per user)
            (generate_ulid_at_time(NOW() - INTERVAL '5 days'), 'Pending Conference 1',
             'Description for pending conference 1', 'Charlie Brown', 'Developer', 'All levels', NULL, 100,
             NOW() + INTERVAL '14 days', NOW() + INTERVAL '14 days' + INTERVAL '2 hours', user1_id, 'pending',
             NOW() - INTERVAL '5 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '5 days'), 'Pending Conference 2',
             'Description for pending conference 2', 'Diana Prince', 'Manager', 'Team leads', 'Management experience',
             50, NOW() + INTERVAL '19 days', NOW() + INTERVAL '19 days' + INTERVAL '2 hours', user2_id, 'pending',
             NOW() - INTERVAL '5 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '5 days'), 'Pending Conference 3',
             'Description for pending conference 3', 'Edward Smith', 'Lead Developer', 'Developers',
             'Coding experience', 75, NOW() + INTERVAL '24 days', NOW() + INTERVAL '24 days' + INTERVAL '2 hours',
             user3_id, 'pending', NOW() - INTERVAL '5 days'),

            -- Rejected conferences
            (generate_ulid_at_time(NOW() - INTERVAL '4 days'), 'Rejected Conference 1',
             'Description for rejected conference 1', 'Frank Miller', 'Developer', 'Beginners', NULL, 100,
             NOW() + INTERVAL '29 days', NOW() + INTERVAL '29 days' + INTERVAL '2 hours', user1_id, 'rejected',
             NOW() - INTERVAL '4 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '4 days'), 'Rejected Conference 2',
             'Description for rejected conference 2', 'Grace Lee', 'Senior Developer', 'Intermediate',
             'Basic programming', 150, NOW() + INTERVAL '33 days', NOW() + INTERVAL '33 days' + INTERVAL '2 hours',
             user2_id, 'rejected', NOW() - INTERVAL '4 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Rejected Conference 3', 'Gaming Development',
             'Quinn Adams', 'Game Developer', 'Game developers', 'C++ knowledge', 90, NOW() + INTERVAL '73 days',
             NOW() + INTERVAL '73 days' + INTERVAL '2 hours', user3_id, 'rejected', NOW() - INTERVAL '3 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Rejected Conference 4', 'Mobile App Security',
             'Rachel Torres', 'Security Engineer', 'Mobile developers', 'iOS/Android development', 80,
             NOW() + INTERVAL '78 days', NOW() + INTERVAL '78 days' + INTERVAL '2 hours', user1_id, 'rejected',
             NOW() - INTERVAL '3 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '3 days'), 'Rejected Conference 5', 'DevOps Best Practices',
             'Sam Lee', 'DevOps Lead', 'Operations teams', 'Linux administration', 100, NOW() + INTERVAL '83 days',
             NOW() + INTERVAL '83 days' + INTERVAL '2 hours', user2_id, 'rejected', NOW() - INTERVAL '3 days'),

            -- Some deleted conferences
            (generate_ulid_at_time(NOW() - INTERVAL '2 days'), 'Deleted Conference 1',
             'Description for deleted conference 1', 'Jack Black', 'Developer', 'All levels', NULL, 100,
             NOW() + INTERVAL '47 days', NOW() + INTERVAL '47 days' + INTERVAL '2 hours', user2_id, 'approved',
             NOW() - INTERVAL '2 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '2 days'), 'Deleted Conference 2',
             'Description for deleted conference 2', 'Kelly White', 'Manager', 'Team leads', 'Management experience',
             75, NOW() + INTERVAL '52 days', NOW() + INTERVAL '52 days' + INTERVAL '2 hours', user3_id, 'pending',
             NOW() - INTERVAL '2 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '2 days'), 'Deleted Conference 3', 'Frontend Testing', 'Tom Wilson',
             'QA Lead', 'Frontend developers', 'JavaScript, React', 70, NOW() + INTERVAL '88 days',
             NOW() + INTERVAL '88 days' + INTERVAL '2 hours', user1_id, 'approved', NOW() - INTERVAL '2 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '2 days'), 'Deleted Conference 4', 'Data Engineering', 'Uma Patel',
             'Data Engineer', 'Data engineers', 'SQL, Python', 85, NOW() + INTERVAL '93 days',
             NOW() + INTERVAL '93 days' + INTERVAL '2 hours', user2_id, 'pending', NOW() - INTERVAL '2 days'),
            (generate_ulid_at_time(NOW() - INTERVAL '2 days'), 'Deleted Conference 5', 'API Design', 'Victor Kim',
             'API Architect', 'Backend developers', 'REST fundamentals', 95, NOW() + INTERVAL '98 days',
             NOW() + INTERVAL '98 days' + INTERVAL '2 hours', user3_id, 'rejected', NOW() - INTERVAL '2 days');


        -- Update deleted conferences
        UPDATE conferences
        SET deleted_at = NOW() - INTERVAL '1 day'
        WHERE title LIKE 'Deleted Conference%';

    END
$$;
