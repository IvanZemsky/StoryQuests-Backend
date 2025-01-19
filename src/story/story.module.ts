import { Module } from "@nestjs/common"
import { StoryController } from "./story.controller"
import { StoryService } from "./story.service"
import { MongooseModule } from "@nestjs/mongoose"
import { Story, StorySchema } from "./story.schema"
import { StoryLike, StoryLikeSchema } from "./storyLike.schema"

@Module({
   imports: [
      MongooseModule.forFeature([
         { name: Story.name, schema: StorySchema },
         { name: StoryLike.name, schema: StoryLikeSchema },
      ]),
   ],
   controllers: [StoryController],
   providers: [StoryService],
})
export class StoryModule {}
