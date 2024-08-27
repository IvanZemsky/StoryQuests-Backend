import { Injectable } from '@nestjs/common';
import { InjectModel } from "@nestjs/mongoose";
import { Scenes } from "./scenes.schema";
import { Model } from "mongoose";

@Injectable()
export class ScenesService {
   constructor(
      @InjectModel(Scenes.name)
      private scenesRepository: Model<Scenes>
   ) {}

   async getScenesByStoryId(storyId: string) {
      const scenes = await this.scenesRepository.find({storyId})
      return scenes
   }
}
