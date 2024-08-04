import torch
from typing import List, Optional
from transformers import AutoModelForSequenceClassification, AutoTokenizer


class CommonRanker:
    def __init__(self, pretrained: str, device: str, batch_size: int = 1):
        self.tokenizer = AutoTokenizer.from_pretrained(pretrained)
        self.model = AutoModelForSequenceClassification.from_pretrained(pretrained)
        self.model.to(device)
        self.model.eval()

    def rank(self, query: str, candidate_list: List[str],
             threshold: Optional[float] = None, top_k: Optional[int] = None) -> List[int]:
        pairs = [(query, candidate) for candidate in candidate_list]
        with torch.no_grad():
            inputs = self.tokenizer(pairs, padding=True, truncation=True, return_tensors='pt', max_length=512)
            scores = self.model(**inputs, return_dict=True).logits.view(-1, ).float()
        sorted_idx_list = sorted(range(len(candidate_list)), key=lambda x: scores[x], reverse=True)
        if threshold:
            sorted_idx_list = list(filter(lambda x: scores[x] >= threshold, sorted_idx_list))
        if top_k:
            sorted_idx_list = sorted_idx_list[:top_k]
        return sorted_idx_list
