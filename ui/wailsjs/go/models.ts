export namespace api {
	export class DirInfo {
		ParentPath: string
		DirName: string
		IsProjectDir: boolean

		static createFrom(source: any = {}) {
			return new DirInfo(source)
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source)
			this.ParentPath = source['ParentPath']
			this.DirName = source['DirName']
			this.IsProjectDir = source['IsProjectDir']
		}
	}
}
