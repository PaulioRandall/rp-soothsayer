import { validateStringField, validateTagObject } from './validateObject'

function testEntity(func, validObject) {
	const invalidId = structuredClone(validObject)
	const invalidName = structuredClone(validObject)

	invalidId.id = 123
	invalidName.name = 123
	expect(() => func(invalidId)).toThrow(Error)
	expect(() => func(invalidName)).toThrow(Error)

	delete invalidId.id
	delete invalidName.name
	expect(() => func(invalidId)).toThrow(Error)
	expect(() => func(invalidName)).toThrow(Error)
}

describe('validateObject.js', () => {
	describe('validateStringField', () => {
		test('good value', () => {
			validateStringField('Thing', 'name', '123')
		})

		test('missing value', () => {
			const func = () => validateStringField('Thing', 'name', null)
			expect(func).toThrow(Error)
		})

		test('bad value', () => {
			const func = () => validateStringField('Thing', 'name', 123)
			expect(func).toThrow(Error)
		})
	})

	describe('validateTagObject()', () => {
		test('no error for good object', () => {
			validateTagObject({
				id: '123',
				name: 'abc',
				description: 'blah',
			})
		})

		test('no error for missing description', () => {
			validateTagObject({
				id: '123',
				name: 'abc',
			})
		})

		test('no error for missing description', () => {
			testEntity(validateTagObject, {
				id: '123',
				name: 'abc',
				description: 'blah',
			})
		})
	})
})
