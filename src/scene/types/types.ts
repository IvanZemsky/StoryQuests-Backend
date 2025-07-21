export type Answer = {
   id: string
   text: string
   nextSceneId: string
}

export type CreateSceneDto = {
   readonly id: string
   readonly testId: string
   readonly answers: Answer[]
   readonly type: "default" | "end"
}