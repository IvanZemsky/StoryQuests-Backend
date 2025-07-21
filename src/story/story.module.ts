import { Module } from "@nestjs/common"
import { StoryController } from "./story.controller"
import { StoryService } from "./story.service"
import { MongooseModule } from "@nestjs/mongoose"
import { Story, StorySchema } from "./schemas/story.schema"
import { StoryLike, StoryLikeSchema } from "./schemas/storyLike.schema"
import { StoryResult, StoryResultSchema } from "./schemas/storyResult.schema"
import { ScenesModule } from "src/scene/scene.module"
import { UserService } from "src/user/user.service"
import { UserModule } from "src/user/user.module"
import { User, UserSchema } from "src/user/user.schema"

@Module({
   imports: [
      MongooseModule.forFeature([
         { name: Story.name, schema: StorySchema },
         { name: StoryLike.name, schema: StoryLikeSchema },
         { name: StoryResult.name, schema: StoryResultSchema },
         { name: User.name, schema: UserSchema },
      ]),
      ScenesModule,
      UserModule,
   ],
   controllers: [StoryController],
   providers: [StoryService, UserService],
})
export class StoryModule {}
