package handlers

import (
	"messengerApp/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SendMessageRequest struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Message    string `json:"message"`
}

type MessageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) SendMessage(c *gin.Context) {
	senderID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sender ID not found"})
		return
	}

	senderIDInt, ok := senderID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid sender ID type"})
		return
	}

	var req struct {
		ReceiverID int    `json:"receiver_id"`
		Message    string `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Проверка, не пытается ли пользователь отправить сообщение самому себе
	if senderIDInt == req.ReceiverID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot send message to self"})
		return
	}

	// Вызов сервиса для отправки сообщения
	if err := h.messageService.SendMessage(senderIDInt, req.ReceiverID, req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	// Получаем параметры из запроса
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	sortBy := c.DefaultQuery("sortBy", "timestamp") // По какому полю сортируем
	sortDir := c.DefaultQuery("sortDir", "desc")    // Направление сортировки: asc или desc
	senderID := c.Query("sender_id")                // Получаем sender_id для фильтрации

	messages, err := h.messageService.GetMessagesWithFilters(page, pageSize, sortBy, sortDir, senderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// Implement other methods as needed
