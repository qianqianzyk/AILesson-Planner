from flask import Flask, request, render_template_string, send_file, abort
import requests  # 用来调用 Moonshot API
from flask_cors import CORS  # 导入 CORS
from flask import Flask, send_file

app = Flask(__name__)

# 启用跨域请求支持
CORS(app)

@app.route('/api/lesson/generate', methods=['POST'])
def generate_lesson_plan():

    print("Received POST request to /api/lesson/generate")

    # 获取 JSON 数据
    data = request.get_json(force=True)
    print("Received JSON:", data)

    if not data:
        return "请求失败: 没有收到数据", 400

    # 获取各个字段
    textbook_name = data.get('textbook_name')
    subject = data.get('subject')
    total_hours = data.get('total_hours')
    topic_name = data.get('topic_name')
    template_file = data.get('template_file')
    resource_file = data.get('resource_file')
    textbook_img = data.get('textbook_img')
    description = data.get('description')

    # 调用 Moonshot API 生成教案内容
    lesson_plan = generate_from_moonshot(textbook_name, subject, total_hours, topic_name, template_file, resource_file, textbook_img, description)

    if "error" in lesson_plan:
        abort(500, description="请求失败: " + lesson_plan["error"])

    # 格式化生成的教案内容
    formatted_lesson_plan = format_lesson_plan(lesson_plan)

    # 生成 markdown 格式的教案内容
    markdown_content = f"""
        {subject} 教案
        教案名称: {textbook_name}
        课题名称: {topic_name}
        总课时: {total_hours}
        教案内容:
        {formatted_lesson_plan}
    """

    # 返回 markdown 格式的教案
    return markdown_content

def generate_from_moonshot(textbook_name, subject, total_hours, topic_name, template_file, resource_file, textbook_img, description):
    """
    调用 Moonshot 大模型 API 生成教案内容
    """
    api_url = "https://api.moonshot.cn/v1/chat/completions"  # 正确的 API 端点
    api_key = ""  # 请使用你的 API 密钥

    prompt_content = (
        f"你是经验丰富的大学讲师，擅长根据学科性质来设计课程。"
        f"请根据以下要求为我生成一份教案：\n\n"
        f"1. 教材名称(请联网搜索,我给你提供的参考资料不包含该教材)：{textbook_name}\n"
        f"2. 学科：{subject}\n"
        f"3. 总课时：{total_hours}\n"
        f"4. 课程主题：{topic_name}\n"
        f"5. 教案模板(请访问Url获取信息)：{template_file}\n"
        f"6. 参考资料(请访问Url获取信息,多个Url以逗号分隔)：{resource_file}\n"
        f"7. 教材图片(请访问Url获取信息,多个Url以逗号分隔)：{textbook_img}\n"
        f"8. 其他要求：{description}\n\n"
        "请严格根据这些信息设计课程内容，要求内容详细完整，课程应包括至少三次师生互动，并考虑到课程的目标和学生的学习需求。"
        "同时为我提供一份完整的教案，包括教学目标、教学内容、互动环节、评估方式等。"
    )   
    
    payload = {
        "model": "moonshot-v1-8k",
        "messages": [{
            "role": "user",
            "content": prompt_content
        }]
    }
    # 设置请求头
    headers = {"Authorization": f"Bearer {api_key}", "Content-Type": "application/json"}

    print("Sending request to Moonshot:", api_url)
    print("Payload:", payload)

    response = requests.post(api_url, json=payload, headers=headers)

    print("Moonshot API Response Status Code:", response.status_code)
    print("Moonshot API Response Content:", response.text)

    if response.status_code == 200:
        return response.json().get('choices')[0].get('message').get('content')
    else:
        return {"error": "Failed to generate lesson plan"}


def format_lesson_plan(lesson_plan):
    """
    格式化教案内容：去掉多余的标记，并进行结构化处理
    """
    # 1. 去除所有的 '*' 标记
    formatted_content = lesson_plan.replace('*', '')

    # 2. 将课程内容根据不同部分分隔（假设每个部分以 '\n' 分隔）
    sections = formatted_content.split('\n')

    # 3. 清理每个部分，去掉每部分前后多余的空格
    cleaned_sections = [section.strip() for section in sections if section.strip()]
    # html_content = markdown.markdown(cleaned_sections, extensions=["extra"]) # 使用 markdown 模块，启用 extra 解析扩展

    # 4. 返回整理后的内容
    return cleaned_sections

if __name__ == "__main__":
    app.run(debug=True, host='0.0.0.0', port=5002)