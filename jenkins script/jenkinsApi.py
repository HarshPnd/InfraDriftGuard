import jenkins


class MyJenkinsAPI:

    def __init__(self):
        self.jenkins_url = 'http://localhost:8080'
        self.server = jenkins.Jenkins('http://localhost:8080', username='Amit', password='password12345')
        self.user = self.server.get_whoami()
        self.version = self.server.get_version()


    def get_user_and_version(self):
        print('Hello %s from Jenkins %s' % (self.user['fullName'], version))

    
    def get_jobs_list(self):
        job_list = []

        for job in self.server.get_jobs():
            full_job_name = job.get('fullname')
            job_list.append(full_job_name)

        return job_list

    def create_job(self):
        pass

    
    def delete_job(self, job_name):
        self.server.delete_job(job_name)


    def build_job(self):
        pass


    def is_job_exists(self, job_name):
        job_list = []

        for job in self.server.get_jobs():
            full_job_name = job.get('fullname')
            job_list.append(full_job_name)
    

        if job_name in job_list:
            return True
        return False

    def get_latest_build(self, full_job_name):
        build_number = self.server.get_job_info(full_job_name) \
            .get("nextBuildNumber") - 1
        
        return build_number

    def get_list_build_numbers(self, full_job_name):
        build_numbers = []
        info = self.server.get_job_info(full_job_name)
        # Loop over builds
        builds = info['builds']

        for build in builds:
            build_number = str(build.get('number'))
            build_numbers.append(build_number)
        
        return build_numbers

    
    def get_console_output(self, full_job_name, job_number):
        return self.server.get_build_console_output(full_job_name, job_number)

    def get_xml_job_config(self, job_name):
        job_xml = self.server.get_job_config(job_name)
        return job_xml

    




api = MyJenkinsAPI()
print(api.get_jobs_list())
# print(api.get_list_build_numbers('A Jenkins Pipeline'))
print(api.get_latest_build('A Jenkins Pipeline'))
# print(api.get_console_output('A Jenkins Pipeline', 10))
print(api.is_job_exists('python-freestyle'))

# Delete a job 
# api.delete_job('react-jenkins')
print(api.get_xml_job_config('A Jenkins Pipeline'))
