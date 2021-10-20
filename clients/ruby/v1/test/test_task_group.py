"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.affinity import Affinity
from openapi_client.model.constraint import Constraint
from openapi_client.model.consul import Consul
from openapi_client.model.ephemeral_disk import EphemeralDisk
from openapi_client.model.migrate_strategy import MigrateStrategy
from openapi_client.model.network_resource import NetworkResource
from openapi_client.model.reschedule_policy import ReschedulePolicy
from openapi_client.model.restart_policy import RestartPolicy
from openapi_client.model.scaling_policy import ScalingPolicy
from openapi_client.model.service import Service
from openapi_client.model.spread import Spread
from openapi_client.model.task import Task
from openapi_client.model.update_strategy import UpdateStrategy
from openapi_client.model.volume_request import VolumeRequest
globals()['Affinity'] = Affinity
globals()['Constraint'] = Constraint
globals()['Consul'] = Consul
globals()['EphemeralDisk'] = EphemeralDisk
globals()['MigrateStrategy'] = MigrateStrategy
globals()['NetworkResource'] = NetworkResource
globals()['ReschedulePolicy'] = ReschedulePolicy
globals()['RestartPolicy'] = RestartPolicy
globals()['ScalingPolicy'] = ScalingPolicy
globals()['Service'] = Service
globals()['Spread'] = Spread
globals()['Task'] = Task
globals()['UpdateStrategy'] = UpdateStrategy
globals()['VolumeRequest'] = VolumeRequest
from openapi_client.model.task_group import TaskGroup


class TestTaskGroup(unittest.TestCase):
    """TaskGroup unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testTaskGroup(self):
        """Test TaskGroup"""
        # FIXME: construct object with mandatory attributes with example values
        # model = TaskGroup()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
