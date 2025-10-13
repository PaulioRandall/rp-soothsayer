import Study from './Study'

describe('models/Study', () => {
	const study = new Study()

	test('has fields and functions', () => {
		// Generic
		expect(typeof study.id).toBe('string')
		expect(typeof study.name).toBe('string')
		expect(typeof study.fromObject).toBe('function')
		expect(typeof study.toObject).toBe('function')

		// Specific
		expect(study.media).toBeInstanceOf(Array)
		expect(study.sessions).toBeInstanceOf(Array)
		expect(study.observationCodes).toBeInstanceOf(Array)
	})
})
