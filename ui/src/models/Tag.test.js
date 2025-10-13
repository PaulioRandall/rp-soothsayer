import Tag from './Tag'

describe('Tag.js', () => {
	test('has fields and functions', () => {
		const tag = new Tag()

		// Generic
		expect(typeof Tag.fromObject).toBe('function')
		expect(typeof tag.id).toBe('string')
		expect(typeof tag.name).toBe('string')
		expect(typeof tag.toObject).toBe('function')

		// Specific
		expect(typeof tag.description).toBe('string')
	})

	describe('Tag.toObject()', () => {
		test('returns expect object', () => {
			const tag = new Tag()
			tag.id = '123'
			tag.name = 'abc'
			tag.description = 'blah'

			expect(tag.toObject()).toEqual({
				id: '123',
				name: 'abc',
				description: 'blah',
			})
		})
	})
})
