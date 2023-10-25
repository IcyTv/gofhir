import asyncio
from fhirpy import AsyncFHIRClient
from fhirpy.base.exceptions import ResourceNotFound
import unittest
import fuzzing

client = AsyncFHIRClient(
    'http://localhost:8080',
    aiohttp_config={'proxy': 'http://localhost:8085'},
)


class ResourceTest(unittest.IsolatedAsyncioTestCase):
    """Abstract class that defines common methods for resource tests"""
    resource: dict
    resource_type: str
    resource_id: str

    async def test_create_resource(self):
        """Test create resource"""
        if "resource_type" not in self.__dict__:
            raise unittest.SkipTest('Resource is not defined')
        resource = client.resource(self.resource_type, **self.resource)
        await resource.save()
        self.assertIsNotNone(resource.id)
        self.resource_id = resource.id


class TestPatient(ResourceTest):
    def __init__(self, t):
        super().__init__(t)
        self.resource_type = 'Patient'
        self.resource = {
            'name': [{
                'given': ['John'],
                'family': 'Smith',
            }],
            'active': True,
        }

    async def test_update_patient(self):
        """Test update patient"""
        resource = client.resource('Patient', **self.resource)
        await resource.save()
        resource.active = False
        await resource.save()
        self.assertFalse(resource.active)


if __name__ == '__main__':
    unittest.defaultTestLoader.loadTestsFromTestCase(TestPatient)

    unittest.main()
