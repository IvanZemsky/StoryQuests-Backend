import { Module } from "@nestjs/common"
import { StoryController } from "./story.controller"
import { StoryService } from "./story.service"
import { MongooseModule } from "@nestjs/mongoose"
import { Story, StorySchema } from "./schemas/story.schema"
import { StoryLike, StoryLikeSchema } from "./schemas/storyLike.schema"
import { StoryResult, StoryResultSchema } from "./schemas/storyResult.schema"
import { ScenesModule } from "src/scene/scene.module"

@Module({
   imports: [
      MongooseModule.forFeature([
         { name: Story.name, schema: StorySchema },
         { name: StoryLike.name, schema: StoryLikeSchema },
         { name: StoryResult.name, schema: StoryResultSchema },
      ]),
      ScenesModule,
   ],
   controllers: [StoryController],
   providers: [StoryService],
})
export class StoryModule {}
