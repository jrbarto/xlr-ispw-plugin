#
# Copyright 2017 XEBIALABS
#
# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
#

import json

from ispw.HttpClient import HttpClient
from ispw.Util import check_response


class AssignmentClient(HttpClient):
    def create_assignment(self, srid, stream, application, default_path, description, owner, assignment_prefix,
                          reference_number, release_id, user_tag):
        context_root = "/ispw/%s/assignments/" % srid
        body = {'stream': stream, 'application': application,
                'defaultPath': default_path,
                'description': description,
                'owner': owner,
                'assignmentPrefix': assignment_prefix,
                'referenceNumber': reference_number,
                'releaseId': release_id,
                'userTag': user_tag}
        response = self._post_request(context_root, json.dumps(body),
                                      {'Accept': 'application/json', 'Content-type': 'application/json'})
        check_response(response, "Failed to create assignment for srid [%s]. Server return [%s], with content [%s]" % (
            srid, response.status_code, response.text))
        print "Called create assignment with id [%s]. Server return [%s], with content [%s]\n" % (
            srid, response.status_code, response.json())
        return response.json()

    def load_task(self, srid, assignment_id, stream, application, module_name, module_type, current_level, starting_level,
                  generate_sequence, sql, ims, cics, program):
        context_root = "/ispw/%s/assignments/%s/tasks" % (srid, assignment_id)
        body = {'application': application, 'stream': stream,
                'moduleName': module_name,
                'moduleType': module_type,
                'currentLevel': current_level,
                'startingLevel': starting_level,
                'generateSequence': generate_sequence,
                'sql': sql,
                'ims': ims,
                'cics': cics,
                'program': program}
        response = self._post_request(context_root, json.dumps(body),
                                      {'Accept': 'application/json', 'Content-type': 'application/json'})
        check_response(response, "Failed to load task for srid [%s]. Server return [%s], with content [%s]" % (
            srid, response.status_code, response.text))
        print "Called load task with id [%s]. Server return [%s], with content [%s]\n" % (
            srid, response.status_code, response.json())
        return response.json()

    def get_assignment_information(self, srid, assignment_id):
        context_root = "/ispw/%s/assignments/%s" % (srid, assignment_id)
        response = self._get_request(context_root,
                                      {'Accept': 'application/json', 'Content-type': 'application/json'})
        check_response(response, "Failed to get assignment information for srid [%s]. Server return [%s], with content [%s]" % (
            srid, response.status_code, response.text))
        print "Called get assignment information with id [%s]. Server return [%s], with content [%s]\n" % (
            srid, response.status_code, response.json())
        return response.json()

    def get_assignment_task_list(self, srid, assignment_id, level):
        context_root = "/ispw/%s/assignments/%s/tasks" % (srid, assignment_id)
        if level:
            context_root += "?level=%s" % level
        response = self._get_request(context_root, {'Accept': 'application/json'})
        check_response(response, "Failed to get assignment task list [%s]. Server return [%s], with content [%s]" % (
            assignment_id, response.status_code, response.text))
        print "Received assignment task list with set id [%s]. Server return [%s], with content [%s]\n" % (
            assignment_id, response.status_code, response.json())
        return response.json()
