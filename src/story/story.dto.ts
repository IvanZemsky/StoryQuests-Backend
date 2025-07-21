import { CreateSceneDto } from "src/scene/types/types"
export type CreateStoryResultDto = {
   readonly resultSceneNumber: string
   readonly datetime: string
}

export type CreateStoryMainInfoDto = {
   readonly name: string
   readonly description: string
   readonly img: string
   readonly tags: string[]
   readonly date: string
}

export type CreateStoryDto = CreateStoryMainInfoDto & {
   readonly scenes: CreateSceneDto[]
}
