import mongo, { MongoClient, ObjectId } from 'mongodb';

const port = process.env.PORT || 8080;

const fhirRequest = (url: string, data?: any) => {
	return fetch(`http://127.0.0.1:${port}/${url}`, {
		method: data ? 'POST' : 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
		body: data ? JSON.stringify(data) : undefined,
	});
};

const mongoUrl = process.env.MONGO_URL || 'mongodb://localhost:27017';
const mongoDbName = process.env.MONGO_DB_NAME || 'fhir';
const client = new MongoClient(mongoUrl);

const testResources: { [key: string]: any[] } = {
	Account: [
		{
			text: {
				status: 'generated',
				div: '<div xmlns="http://www.w3.org/1999/xhtml">HACC Funded Billing for Peter James Chalmers</div>',
			},
			identifier: [
				{
					system: 'urn:oid:0.1.2.3.4.5.6.7',
					value: '654321',
				},
			],
			status: 'active',
			type: {
				coding: [
					{
						system: 'http://terminology.hl7.org/CodeSystem/v3-ActCode',
						code: 'PBILLACCT',
						display: 'patient billing account',
					},
				],
				text: 'patient',
			},
			name: 'HACC Funded Billing for Peter James Chalmers',
			subject: [
				{
					reference: 'Patient/example',
					display: 'Peter James Chalmers',
				},
			],
			owner: {
				reference: 'Organization/hl7',
			},
			description: 'Hospital charges',
		},
	],
};

const testResourceIds: { [key: string]: string[] } = {};

describe('fetch', () => {
	beforeAll(async () => {
		await client.connect();
		const db = client.db(mongoDbName);

		for (const resourceType in testResources) {
			const collection = db.collection(resourceType);
			const res = await collection.insertMany(
				testResources[resourceType]
			);

			// Update ids
			testResourceIds[resourceType] = Object.values(res.insertedIds).map(
				(id) => id.toString()
			);
		}
	});

	afterAll(async () => {
		// Remove all test data
		const db = client.db(mongoDbName);
		for (const resourceType in testResourceIds) {
			const collection = db.collection(resourceType);
			for (const id of testResourceIds[resourceType]) {
				await collection.deleteOne({ _id: new ObjectId(id) });
			}
		}

		await client.close();
	});

	it('should fetch a single Account', async () => {
		const id = testResourceIds.Account[0];
		const response = await fhirRequest(`Account/${id}`);
		expect(response.status).toBe(200);
		const data = await response.json();
		expect(data.resourceType).toBe('Account');
		expect(data.id).toBe(id);

		// Compare response to test data
		const testResource = testResources.Account[0];
		testResource.id = id;

		expect(data).toEqual(testResource);
	});
});
