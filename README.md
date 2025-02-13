### 1、本地ollama的API接口

接着上一章所[本地部署deepseek](https://www.codezhou.top/article/%E6%9C%AC%E5%9C%B0%E9%83%A8%E7%BD%B2deepseek)，这一章我们调用ollama api

![image-20250213145757287](https://codegym.oss-cn-shenzhen.aliyuncs.com/uiiujhj/202502131457049.png)



对应的curl：

```bash
curl --request POST \
  --url http://localhost:11434/api/generate \
  --header 'Accept: */*' \
  --header 'Accept-Encoding: gzip, deflate, br' \
  --header 'Connection: keep-alive' \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: PostmanRuntime-ApipostRuntime/1.1.0' \
  --data '{"model":"deepseek-r1:7b","prompt":"你是谁","stream":false}1'
```



### 2、java版本的API调用

java版本对应的github地址

```bash
https://github.com/cativen/chat-api-java/
```

pom.xml

```java
  <dependencies>
        <dependency>
            <groupId>com.squareup.okhttp3</groupId>
            <artifactId>okhttp-sse</artifactId>
            <version>3.14.9</version>
        </dependency>
        <dependency>
            <groupId>com.squareup.okhttp3</groupId>
            <artifactId>logging-interceptor</artifactId>
            <version>3.14.9</version>
        </dependency>
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>fastjson</artifactId>
            <version>1.2.83</version> <!-- 请检查并使用最新版本 -->
        </dependency>
    </dependencies>
```



ChatRequest

```java
public class ChatRequest {
    private String model;

    private String prompt;

    private Boolean stream;

    public String getModel() {
        return model;
    }

    public void setModel(String model) {
        this.model = model;
    }

    public String getPrompt() {
        return prompt;
    }

    public void setPrompt(String prompt) {
        this.prompt = prompt;
    }

    public Boolean getStream() {
        return stream;
    }

    public void setStream(Boolean stream) {
        this.stream = stream;
    }
}

```



ChatResponse

```java
public class ChatResponse {

    private String model;

    private String response;

    private String created_at;

    public String getModel() {
        return model;
    }

    public void setModel(String model) {
        this.model = model;
    }

    public String getResponse() {
        return response;
    }

    public void setResponse(String response) {
        this.response = response;
    }

    public String getCreated_at() {
        return created_at;
    }

    public void setCreated_at(String created_at) {
        this.created_at = created_at;
    }
}
```



ApiTest

```java
import com.alibaba.fastjson.JSON;
import okhttp3.*;

import java.io.IOException;
import java.util.concurrent.TimeUnit;

public class ApiTest {
    public static void main(String[] args) {
        // 2. 开启 Http 客户端
        OkHttpClient client = new OkHttpClient
                .Builder()
                .connectTimeout(50, TimeUnit.SECONDS)
                .writeTimeout(50, TimeUnit.SECONDS)
                .readTimeout(50, TimeUnit.SECONDS)
                .build();

        // 构建 JSON 请求体
        ChatRequest chatRequest = new ChatRequest();
        chatRequest.setStream(false);
        chatRequest.setModel("deepseek-r1:7b");
        chatRequest.setPrompt("如何学英语");
        String jsonString = JSON.toJSONString(chatRequest);
        MediaType jsonType = MediaType.get("application/json; charset=utf-8");
        RequestBody body = RequestBody.create(jsonType,jsonString);

        // 创建 POST 请求
        Request request = new Request.Builder()
                .url("http://localhost:11434/api/generate")
                .addHeader("Content-Type", "application/json")
                .post(body)
                .build();

        // 发送同步 POST 请求
        try (Response response = client.newCall(request).execute()) {
            if (response.isSuccessful()) {
                ChatResponse chatResponse = JSON.parseObject(response.body().string(), ChatResponse.class);
                System.out.println(chatResponse.getResponse());
            } else {
                System.err.println("Request failed: " + response.code());
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

结果：

```bash
学习英语可以分为以下几个步骤：

### 1. **打牢基础**
   - **听写训练**：每天听英文歌曲、短文或视频，试着跟读并写下听到的单词和句子。
   - **背诵单词**：使用 flashcards（正反面打印）来记忆词汇。可以用 apps 或 websites 如 Anki, Quizlet 等。
   - **掌握发音**：学习英语字母表、元音和辅音的发音，区分重音符号（`'`）。多听录音，模仿发音。

### 2. **提高听力**
   - **多听音频**：每天花时间听英文新闻、 talk shows 或音乐。可以从 BBC 学习 English, Duolingo Radio 等渠道获取内容。
   - **观看电影/视频**：选择英语水平相近的电影或教学视频，尝试理解对话和解释。

### 3. **增强阅读能力**
   - **读简单文章**：从基本的 grammar 到新闻报道开始。可以使用 apps 如 ESL Kids 或 Readwise。
   - **学习语法**：通过教材或在线课程（如 Khan Academy, Duolingo）系统地学习 grammar，确保理解每个规则。

### 4. **练习口语**
   - **和外教对话**：如果可能的话，找一个会英语的伴练。如果没有，可以用 apps 或 tools 如 Tandem 进行线上配对。
   - **录播回放**：多录音并回看，找出发音、语法或语调上的错误，加以改进。

### 5. **提高写作能力**
   - **写日记或短文**：每天写几句话或一篇小文章，记录生活或想法。可以使用 apps 如 Grammarly 来检查错误。
   - **模仿优秀文章**：分析 High-Commissioner 的文章，学习句式结构和表达方式。

### 6. **参加活动**
   - **加入英语club**：与英语学习者交流，参加讨论会或比赛，提升自信心。
   - **参与写作比赛**：通过 Writing Contests 等平台提交作品，得到反馈并改进。

### 7. **保持耐心和坚持**
   - **每天练习**：英语的学习需要持续性，每天花一定时间练习听、说、读、写。
   - **设定目标**：制定短期和长期学习目标，并逐步实现，保持动力。

### 8. **利用技术工具**
   - **使用 apps 和工具**：如 Duolingo, HelloChinese, Memrise 等来辅助学习。
   - **在线课程**：如果需要系统化的学习，可以选择 online courses（如 Coursera, edX）或参加英语学校。

通过以上步骤的系统学习和不断的实践，可以逐步提高英语水平。记得保持积极心态，英语学习是一个长期的过程。

Process finished with exit code 0

```





### 3、go版本的API调用

go版本对应的github地址

```bash
https://github.com/cativen/chat-api-go
```

mian.go

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Completion struct {
	Response  string `json:"response"`
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
}

type ChatRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

func main() {
	url := "http://localhost:11434/api/generate"
	// 创建请求体结构体
	requestBody := ChatRequest{
		Model:  "deepseek-r1:7b",
		Prompt: "如何做好家里的卫生工作",
		Stream: false,
	}

	// 将结构体转换为 JSON
	payload, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer sk-b8ebb99508964850b2b1c")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var completion Completion
	err = json.Unmarshal(body, &completion)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Content:", completion.Response)
}
```



结果

```bash
做好家里的卫生工作是每个家庭成员共同努力的结果。以下是一些实用的建议和步骤，可以帮助您改善家庭卫生，创造一个干净、舒适的生活环境：

### 1. **划分责任区**
   - 家庭成员可以轮流负责不同的区域（如客厅、厨房、卧室等），明确每个人的卫生责任。
   - 这样可以让大家感受到责任感，并更容易监督和检查自己的区域是否清洁。

### 2. **保持日常清洁**
   - **每天打扫**：花几分钟时间整理桌面、书架或其他需要清洁的地方，保持整洁。
   - **每周性任务**：
     - 垃圾分类：将垃圾区的垃圾归类并丢弃到相应的垃圾桶。
     - 垃圾处理：定期清理厨房和卫生间，减少积存。
     - 环境维护：每天早晨或晚上花5-10分钟扫地，保持地面干净。

### 3. **日常清洁**
   - 地面：使用地拖或吸尘器定期清扫地板，特别是厨房、餐厅和卧室。
   - 卫生间和厨房：每天用清水擦洗台面、镜柜和下水道，定期清洁卫生间的马桶。
   - 每周一次彻底打扫：清理沙发、窗帘、床单等布艺品，拖地或用湿布擦拭墙面。

### 4. **注意细节**
   - 家里的角落、边缘、楼梯扶手等地方容易被忽视，需要特别关注。
   - 厨房和卫生间是常见的污染源，及时清理和清洁可以减少异味。
   - 使用温和的清洁剂或中性洗涤剂清洗家具表面，避免留下 streaks。

### 5. **培养良好习惯**
   - 手段关灯：养成随手关闭 lights 的习惯，减少灰尘积累。
   - 不随便丢垃圾：将垃圾分类后放入指定垃圾桶，避免长时间暴露在空气中。
   - 定期大扫除：每周进行一次彻底的打扫，清理所有区域，包括地面、墙面和家具。

### 6. **使用适当的工具**
   - 地拖：拖地时选择轻便且耐用的地拖，可以有效清洁地面。
   - 吸尘器：吸尘器可以帮助去除地毯和其他表面的灰尘。
   - 墨水：对于墙面或家具上的污渍，可以用温和的墨水擦拭。

### 7. **互相监督**
   - 家人可以轮流检查对方是否完成卫生任务，比如“检查一下厨房的桌子是否整洁”或者“看看卫生间有没有倒掉马桶里的水”。
   - 如果发现卫生状况下降，及时沟通，共同寻找原因并解决。

通过以上步骤和习惯的养成，您将逐渐改善家里的卫生状况。记住，保持卫生不仅是为了个人健康，也是为了家庭环境的整体舒适。
```



