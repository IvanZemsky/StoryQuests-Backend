import { Injectable } from "@nestjs/common"
import { InjectModel } from "@nestjs/mongoose"
import { Scene } from "./scene.schema"
import mongoose, { Model } from "mongoose"
import { CreateSceneDto } from "./types/types"

@Injectable()
export class SceneService {
   constructor(
      @InjectModel(Scene.name)
      private sceneModel: Model<Scene>,
   ) {}

   async getScenesByStoryId(storyId: string) {
      const scenes = await this.sceneModel.find({ storyId })
      return scenes
   }

   async getScene(searchParams: { storyId?: string; number?: string }) {
      const scene = await this.sceneModel.findOne(searchParams).lean()
      return scene
   }

   async getEndScenes(storyId: string) {
      const scenes = await this.sceneModel.find({ storyId, type: "end" })
      return scenes
   }

   async createScenes(storyId: string, scenes: CreateSceneDto[]) {
      const scenesToCreate = scenes.map((scene) => {
         const sceneObj: CreateSceneDto & {
            passes?: number
            storyId: mongoose.Types.ObjectId
         } = {
            ...scene,
            storyId: new mongoose.Types.ObjectId(storyId),
         }
         if (scene.type === "end") {
            sceneObj.passes = 0
         }
         return sceneObj
      })
      const newScenes = await this.sceneModel.create(scenesToCreate)
      return newScenes
   }

   async incrementPasses(storyId: string, sceneNumber: string) {
      const scene = await this.sceneModel.findOneAndUpdate(
         { storyId, number: sceneNumber },
         { $inc: { passes: 1 } },
         { new: true },
      )
      return scene
   }
}
