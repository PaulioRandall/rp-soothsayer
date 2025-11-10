export namespace api {
	export class DirInfo {
		DirName: string
		DirPath: string
		IsProjectDir: boolean

		static createFrom(source: any = {}) {
			return new DirInfo(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.DirName = source['DirName']
			this.DirPath = source['DirPath']
			this.IsProjectDir = source['IsProjectDir']
		}
	}
}

export namespace study {
	export class Code {
		id: string
		name: string
		tags: string
		description: string

		static createFrom(source: any = {}) {
			return new Code(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.id = source['id']
			this.name = source['name']
			this.tags = source['tags']
			this.description = source['description']
		}
	}
	export class Observation {
		id: string
		name: string
		tags: string
		start: number
		quote: string
		description: string
		codes: string[]

		static createFrom(source: any = {}) {
			return new Observation(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.id = source['id']
			this.name = source['name']
			this.tags = source['tags']
			this.start = source['start']
			this.quote = source['quote']
			this.description = source['description']
			this.codes = source['codes']
		}
	}
	export class MediaSection {
		id: string
		name: string
		tags: string
		start: number
		end: number

		static createFrom(source: any = {}) {
			return new MediaSection(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.id = source['id']
			this.name = source['name']
			this.tags = source['tags']
			this.start = source['start']
			this.end = source['end']
		}
	}
	export class Media {
		id: string
		name: string
		tags: string
		filepath: string
		sections: MediaSection[]
		observations: Observation[]

		static createFrom(source: any = {}) {
			return new Media(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.id = source['id']
			this.name = source['name']
			this.tags = source['tags']
			this.filepath = source['filepath']
			this.sections = this.convertValues(source['sections'], MediaSection)
			this.observations = this.convertValues(
				source['observations'],
				Observation
			)
		}

		convertValues(a: any, classs: any, asMap: boolean = false): any {
			if (!a) {
				return a
			}
			if (a.slice && a.map) {
				return (a as any[]).map((elem) => this.convertValues(elem, classs))
			} else if ('object' === typeof a) {
				if (asMap) {
					for (const key of Object.keys(a)) {
						a[key] = new classs(a[key])
					}
					return a
				}
				return new classs(a)
			}
			return a
		}
	}

	export class Tag {
		id: string
		name: string
		description: string

		static createFrom(source: any = {}) {
			return new Tag(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.id = source['id']
			this.name = source['name']
			this.description = source['description']
		}
	}
	export class Study {
		name: string
		filepath: string
		media: Media[]
		code: Code[]
		tag: Tag[]

		static createFrom(source: any = {}) {
			return new Study(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.name = source['name']
			this.filepath = source['filepath']
			this.media = this.convertValues(source['media'], Media)
			this.code = this.convertValues(source['code'], Code)
			this.tag = this.convertValues(source['tag'], Tag)
		}

		convertValues(a: any, classs: any, asMap: boolean = false): any {
			if (!a) {
				return a
			}
			if (a.slice && a.map) {
				return (a as any[]).map((elem) => this.convertValues(elem, classs))
			} else if ('object' === typeof a) {
				if (asMap) {
					for (const key of Object.keys(a)) {
						a[key] = new classs(a[key])
					}
					return a
				}
				return new classs(a)
			}
			return a
		}
	}
}
