export default function globalSetup() {
	globalThis.g_or = g_or
	globalThis.g_toObject = g_toObject
	globalThis.g_toObjects = g_toObjects
	globalThis.g_populateArrayFromObjectArray = g_populateArrayFromObjectArray
	globalThis.g_validateStringField = g_validateStringField
}

function g_or(v, other) {
	return v ? v : other
}

function g_toObject(v) {
	return v.toObject()
}

function g_toObjects(array) {
	return array.map(g_toObject)
}

function g_populateArrayFromObjectArray(array, objectArray, fromObjectFunc) {
	for (const obj of objectArray) {
		const v = fromObjectFunc(obj)
		array.push(v)
	}
}

function g_validateStringField(objType, fieldName, value) {
	if (typeof value !== 'string') {
		throw new Error(`${objType} ${fieldName} must be a string`)
	}

	if (value.trim() === '') {
		throw new Error(`${objType} ${fieldName} must be a non-empty`)
	}
}
