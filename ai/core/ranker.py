import torch
from typing import List
from transformers import AutoModelForSequenceClassification, AutoTokenizer


class CommonRanker:
    def __init__(self, pretrained: str, device: str):
        self.tokenizer = AutoTokenizer.from_pretrained(pretrained)
        self.model = AutoModelForSequenceClassification.from_pretrained(pretrained)
        self.model.eval()

    def rank(self, query: str, candidate_list: List[str]) -> List[str]:
        pairs = [(query, candidate) for candidate in candidate_list]
        with torch.no_grad():
            inputs = self.tokenizer(pairs, padding=True, truncation=True, return_tensors='pt', max_length=512)
            scores = self.model(**inputs, return_dict=True).logits.view(-1, ).float()
