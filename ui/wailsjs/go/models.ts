export namespace api {
	export class DirInfo {
		DirName: string
		IsProjectDir: boolean

		static createFrom(source: any = {}) {
			return new DirInfo(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.DirName = source['DirName']
			this.IsProjectDir = source['IsProjectDir']
		}
	}
}
