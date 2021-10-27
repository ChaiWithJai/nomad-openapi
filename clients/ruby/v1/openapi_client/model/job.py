"""
    Nomad

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 1.1.4
    Contact: support@hashicorp.com
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from openapi_client.model_utils import (  # noqa: F401
    ApiTypeError,
    ModelComposed,
    ModelNormal,
    ModelSimple,
    cached_property,
    change_keys_js_to_python,
    convert_js_args_to_python_args,
    date,
    datetime,
    file_type,
    none_type,
    validate_get_composed_info,
)
from ..model_utils import OpenApiModel
from openapi_client.exceptions import ApiAttributeError


def lazy_import():
    from openapi_client.model.affinity import Affinity
    from openapi_client.model.constraint import Constraint
    from openapi_client.model.migrate_strategy import MigrateStrategy
    from openapi_client.model.multiregion import Multiregion
    from openapi_client.model.parameterized_job_config import ParameterizedJobConfig
    from openapi_client.model.periodic_config import PeriodicConfig
    from openapi_client.model.reschedule_policy import ReschedulePolicy
    from openapi_client.model.spread import Spread
    from openapi_client.model.task_group import TaskGroup
    from openapi_client.model.update_strategy import UpdateStrategy
    globals()['Affinity'] = Affinity
    globals()['Constraint'] = Constraint
    globals()['MigrateStrategy'] = MigrateStrategy
    globals()['Multiregion'] = Multiregion
    globals()['ParameterizedJobConfig'] = ParameterizedJobConfig
    globals()['PeriodicConfig'] = PeriodicConfig
    globals()['ReschedulePolicy'] = ReschedulePolicy
    globals()['Spread'] = Spread
    globals()['TaskGroup'] = TaskGroup
    globals()['UpdateStrategy'] = UpdateStrategy


class Job(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
    }

    validations = {
        ('create_index',): {
            'inclusive_maximum': 384,
            'inclusive_minimum': 0,
        },
        ('job_modify_index',): {
            'inclusive_maximum': 384,
            'inclusive_minimum': 0,
        },
        ('modify_index',): {
            'inclusive_maximum': 384,
            'inclusive_minimum': 0,
        },
        ('version',): {
            'inclusive_maximum': 384,
            'inclusive_minimum': 0,
        },
    }

    @cached_property
    def additional_properties_type():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded
        """
        lazy_import()
        return (bool, date, datetime, dict, float, int, list, str, none_type,)  # noqa: E501

    _nullable = False

    @cached_property
    def openapi_types():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded

        Returns
            openapi_types (dict): The key is attribute name
                and the value is attribute type.
        """
        lazy_import()
        return {
            'affinities': ([Affinity],),  # noqa: E501
            'all_at_once': (bool,),  # noqa: E501
            'constraints': ([Constraint],),  # noqa: E501
            'consul_namespace': (str,),  # noqa: E501
            'consul_token': (str,),  # noqa: E501
            'create_index': (int,),  # noqa: E501
            'datacenters': ([str],),  # noqa: E501
            'dispatched': (bool,),  # noqa: E501
            'id': (str,),  # noqa: E501
            'job_modify_index': (int,),  # noqa: E501
            'meta': ({str: (str,)},),  # noqa: E501
            'migrate': (MigrateStrategy,),  # noqa: E501
            'modify_index': (int,),  # noqa: E501
            'multiregion': (Multiregion,),  # noqa: E501
            'name': (str,),  # noqa: E501
            'namespace': (str,),  # noqa: E501
            'nomad_token_id': (str,),  # noqa: E501
            'parameterized_job': (ParameterizedJobConfig,),  # noqa: E501
            'parent_id': (str,),  # noqa: E501
            'payload': (str,),  # noqa: E501
            'periodic': (PeriodicConfig,),  # noqa: E501
            'priority': (int,),  # noqa: E501
            'region': (str,),  # noqa: E501
            'reschedule': (ReschedulePolicy,),  # noqa: E501
            'spreads': ([Spread],),  # noqa: E501
            'stable': (bool,),  # noqa: E501
            'status': (str,),  # noqa: E501
            'status_description': (str,),  # noqa: E501
            'stop': (bool,),  # noqa: E501
            'submit_time': (int,),  # noqa: E501
            'task_groups': ([TaskGroup],),  # noqa: E501
            'type': (str,),  # noqa: E501
            'update': (UpdateStrategy,),  # noqa: E501
            'vault_namespace': (str,),  # noqa: E501
            'vault_token': (str,),  # noqa: E501
            'version': (int,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'affinities': 'Affinities',  # noqa: E501
        'all_at_once': 'AllAtOnce',  # noqa: E501
        'constraints': 'Constraints',  # noqa: E501
        'consul_namespace': 'ConsulNamespace',  # noqa: E501
        'consul_token': 'ConsulToken',  # noqa: E501
        'create_index': 'CreateIndex',  # noqa: E501
        'datacenters': 'Datacenters',  # noqa: E501
        'dispatched': 'Dispatched',  # noqa: E501
        'id': 'ID',  # noqa: E501
        'job_modify_index': 'JobModifyIndex',  # noqa: E501
        'meta': 'Meta',  # noqa: E501
        'migrate': 'Migrate',  # noqa: E501
        'modify_index': 'ModifyIndex',  # noqa: E501
        'multiregion': 'Multiregion',  # noqa: E501
        'name': 'Name',  # noqa: E501
        'namespace': 'Namespace',  # noqa: E501
        'nomad_token_id': 'NomadTokenID',  # noqa: E501
        'parameterized_job': 'ParameterizedJob',  # noqa: E501
        'parent_id': 'ParentID',  # noqa: E501
        'payload': 'Payload',  # noqa: E501
        'periodic': 'Periodic',  # noqa: E501
        'priority': 'Priority',  # noqa: E501
        'region': 'Region',  # noqa: E501
        'reschedule': 'Reschedule',  # noqa: E501
        'spreads': 'Spreads',  # noqa: E501
        'stable': 'Stable',  # noqa: E501
        'status': 'Status',  # noqa: E501
        'status_description': 'StatusDescription',  # noqa: E501
        'stop': 'Stop',  # noqa: E501
        'submit_time': 'SubmitTime',  # noqa: E501
        'task_groups': 'TaskGroups',  # noqa: E501
        'type': 'Type',  # noqa: E501
        'update': 'Update',  # noqa: E501
        'vault_namespace': 'VaultNamespace',  # noqa: E501
        'vault_token': 'VaultToken',  # noqa: E501
        'version': 'Version',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """Job - a model defined in OpenAPI

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            affinities ([Affinity]): [optional]  # noqa: E501
            all_at_once (bool): [optional]  # noqa: E501
            constraints ([Constraint]): [optional]  # noqa: E501
            consul_namespace (str): [optional]  # noqa: E501
            consul_token (str): [optional]  # noqa: E501
            create_index (int): [optional]  # noqa: E501
            datacenters ([str]): [optional]  # noqa: E501
            dispatched (bool): [optional]  # noqa: E501
            id (str): [optional]  # noqa: E501
            job_modify_index (int): [optional]  # noqa: E501
            meta ({str: (str,)}): [optional]  # noqa: E501
            migrate (MigrateStrategy): [optional]  # noqa: E501
            modify_index (int): [optional]  # noqa: E501
            multiregion (Multiregion): [optional]  # noqa: E501
            name (str): [optional]  # noqa: E501
            namespace (str): [optional]  # noqa: E501
            nomad_token_id (str): [optional]  # noqa: E501
            parameterized_job (ParameterizedJobConfig): [optional]  # noqa: E501
            parent_id (str): [optional]  # noqa: E501
            payload (str): [optional]  # noqa: E501
            periodic (PeriodicConfig): [optional]  # noqa: E501
            priority (int): [optional]  # noqa: E501
            region (str): [optional]  # noqa: E501
            reschedule (ReschedulePolicy): [optional]  # noqa: E501
            spreads ([Spread]): [optional]  # noqa: E501
            stable (bool): [optional]  # noqa: E501
            status (str): [optional]  # noqa: E501
            status_description (str): [optional]  # noqa: E501
            stop (bool): [optional]  # noqa: E501
            submit_time (int): [optional]  # noqa: E501
            task_groups ([TaskGroup]): [optional]  # noqa: E501
            type (str): [optional]  # noqa: E501
            update (UpdateStrategy): [optional]  # noqa: E501
            vault_namespace (str): [optional]  # noqa: E501
            vault_token (str): [optional]  # noqa: E501
            version (int): [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        self = super(OpenApiModel, cls).__new__(cls)

        if args:
            raise ApiTypeError(
                "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                    args,
                    self.__class__.__name__,
                ),
                path_to_item=_path_to_item,
                valid_classes=(self.__class__,),
            )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
        return self

    required_properties = set([
        '_data_store',
        '_check_type',
        '_spec_property_naming',
        '_path_to_item',
        '_configuration',
        '_visited_composed_classes',
    ])

    @convert_js_args_to_python_args
    def __init__(self, *args, **kwargs):  # noqa: E501
        """Job - a model defined in OpenAPI

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            affinities ([Affinity]): [optional]  # noqa: E501
            all_at_once (bool): [optional]  # noqa: E501
            constraints ([Constraint]): [optional]  # noqa: E501
            consul_namespace (str): [optional]  # noqa: E501
            consul_token (str): [optional]  # noqa: E501
            create_index (int): [optional]  # noqa: E501
            datacenters ([str]): [optional]  # noqa: E501
            dispatched (bool): [optional]  # noqa: E501
            id (str): [optional]  # noqa: E501
            job_modify_index (int): [optional]  # noqa: E501
            meta ({str: (str,)}): [optional]  # noqa: E501
            migrate (MigrateStrategy): [optional]  # noqa: E501
            modify_index (int): [optional]  # noqa: E501
            multiregion (Multiregion): [optional]  # noqa: E501
            name (str): [optional]  # noqa: E501
            namespace (str): [optional]  # noqa: E501
            nomad_token_id (str): [optional]  # noqa: E501
            parameterized_job (ParameterizedJobConfig): [optional]  # noqa: E501
            parent_id (str): [optional]  # noqa: E501
            payload (str): [optional]  # noqa: E501
            periodic (PeriodicConfig): [optional]  # noqa: E501
            priority (int): [optional]  # noqa: E501
            region (str): [optional]  # noqa: E501
            reschedule (ReschedulePolicy): [optional]  # noqa: E501
            spreads ([Spread]): [optional]  # noqa: E501
            stable (bool): [optional]  # noqa: E501
            status (str): [optional]  # noqa: E501
            status_description (str): [optional]  # noqa: E501
            stop (bool): [optional]  # noqa: E501
            submit_time (int): [optional]  # noqa: E501
            task_groups ([TaskGroup]): [optional]  # noqa: E501
            type (str): [optional]  # noqa: E501
            update (UpdateStrategy): [optional]  # noqa: E501
            vault_namespace (str): [optional]  # noqa: E501
            vault_token (str): [optional]  # noqa: E501
            version (int): [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        if args:
            raise ApiTypeError(
                "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                    args,
                    self.__class__.__name__,
                ),
                path_to_item=_path_to_item,
                valid_classes=(self.__class__,),
            )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")