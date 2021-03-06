#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from unittest import TestCase
from unittest.mock import MagicMock

from polycommon.unique_urls import (
    get_project_url,
    get_run_health_url,
    get_run_url,
    get_user_project_url,
    get_user_url,
)


class TestUniqueUrls(TestCase):
    def setUp(self):
        super().setUp()
        self.user = MagicMock(username="foo")
        unique_name = f"{self.user.username}.project-test"
        self.project = MagicMock(
            id=1, name="project-test", user=self.user, unique_name=unique_name
        )
        unique_name = f"{unique_name}.uuid"
        self.run = MagicMock(
            uuid="uuid", id=2, project=self.project, unique_name=unique_name
        )

    def test_get_user_url(self):
        self.assertEqual(get_user_url("foo"), "/foo")

    def test_get_project_url(self):
        unique_name = self.project.unique_name
        self.assertEqual(
            get_project_url(unique_name=unique_name),
            "/{}".format(unique_name.replace(".", "/")),
        )

    def test_get_user_project_url(self):
        self.assertEqual(
            get_user_project_url(username="foo", project_name="bar"), "/foo/bar"
        )

    def test_get_run_url(self):
        unique_name = self.run.unique_name
        self.assertEqual(
            get_run_url(unique_name=unique_name), "/foo/project-test/runs/uuid"
        )

    def test_get_run_health_url(self):
        self.assertEqual(
            get_run_health_url(unique_name=self.run.unique_name),
            "{}/_heartbeat".format(get_run_url(self.run.unique_name)),
        )
